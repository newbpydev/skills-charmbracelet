import { readFile, writeFile, mkdir, readdir, mkdtemp, rm } from 'node:fs/promises';
import { run as runCommand } from './lib.mjs';
import { join, resolve } from 'node:path';
import { tmpdir } from 'node:os';
import { fileURLToPath } from 'node:url';

const root = fileURLToPath(new URL('../', import.meta.url));
const [id, variant = 'baseline', checkOnly] = process.argv.slice(2);
const cases = JSON.parse(await readFile(join(root, 'evaluations/cases.json'), 'utf8'));
const task = cases.find(c => c.id === id);
if (!task || !['baseline', 'skill'].includes(variant)) {
  throw new Error('Usage: node scripts/evaluate.mjs <new-v2|legacy-fix|migrate-v2> <baseline|skill> [--check-only]');
}
const dir = join(root, '.cache/evaluations', `${id}-${variant}`);
await mkdir(dir, { recursive: true });
const mod = `module evaluation\n\ngo 1.25.12\n\nrequire (\n${Object.entries(task.modules).map(([m, v]) => `\t${m} ${v}`).join('\n')}\n)\n`;
const run = (command, args, cwd, input = '', timeout = 180_000) => runCommand(command, args, { cwd, input, timeout });

if (checkOnly !== '--check-only') {
  let prompt = `${task.prompt}\n\nTarget go.mod:\n${mod}`;
  if (task.seed) prompt += `\nExisting source:\n${await readFile(join(root, 'evaluations', task.seed), 'utf8')}`;
  const suppliedFiles = [];
  if (variant === 'skill') {
    for (const name of task.skills) {
      const base = join(root, 'skills', `charmbracelet-${name}`);
      for (const file of ['SKILL.md', ...(await readdir(join(base, 'references'))).filter(n => n.endsWith('.md')).map(n => `references/${n}`)]) {
        suppliedFiles.push(`skills/charmbracelet-${name}/${file}`);
        prompt += `\n\nReference ${name}/${file}:\n${await readFile(join(base, file), 'utf8')}`;
      }
    }
  }
  prompt += '\nReturn ONLY a JSON object with keys code (the complete main.go as a string) and notes (a brief string). No markdown fences.';
  await writeFile(join(dir, 'prompt.txt'), prompt);
  const scratch = await mkdtemp(join(tmpdir(), 'charm-eval-'));
  const started = Date.now();
  try {
    const schemaPath = join(scratch, 'output.schema.json');
    await writeFile(schemaPath, JSON.stringify({ type: 'object', properties: { code: { type: 'string' }, notes: { type: 'string' } }, required: ['code', 'notes'], additionalProperties: false }));
    const result = await run('codex', ['exec', '--ignore-user-config', '--ephemeral', '--skip-git-repo-check', '--sandbox', 'read-only', '--color', 'never', '--json', '--output-schema', schemaPath, '-o', join(dir, 'answer.json'), '-C', scratch, '-'], scratch, 'Complete this bounded code-generation evaluation using only the supplied context. Do not invoke any tools or read other files.\n\n' + prompt);
    await writeFile(join(dir, 'response.jsonl'), result.stdout);
    await writeFile(join(dir, 'execution.json'), JSON.stringify({ command: 'codex', version: (await run('codex', ['--version'], scratch)).stdout.trim(), exitCode: result.code, signal: result.signal, elapsedMs: Date.now() - started, suppliedFiles, stderr: result.stderr }, null, 2) + '\n');
    if (result.code !== 0) throw new Error(`Evaluation provider failed; see ${dir}/execution.json`);
  } finally { await rm(scratch, { recursive: true, force: true }); }
}
const output = JSON.parse(await readFile(join(dir, 'answer.json'), 'utf8'));
if (typeof output.code !== 'string') throw new Error('Provider did not return code');
await writeFile(join(dir, 'main.go'), output.code);
await writeFile(join(dir, 'go.mod'), mod);
await writeFile(join(dir, 'notes.txt'), String(output.notes ?? ''));
const tidy = await run('go', ['mod', 'tidy'], dir);
const build = tidy.code === 0 ? await run('go', ['test', '-mod=readonly', './...'], dir) : null;
await writeFile(join(dir, 'check.json'), JSON.stringify({ tidy, build, criteria: task.criteria }, null, 2) + '\n');
console.log(JSON.stringify({ id, variant, directory: resolve(dir), compiled: build?.code === 0, output: build?.stdout || tidy.stderr }));
if (build?.code !== 0) process.exitCode = 1;
