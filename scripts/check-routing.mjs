import { mkdtemp, mkdir, readFile, readdir, writeFile, rm } from 'node:fs/promises';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { parse } from 'yaml';
import { root, run } from './lib.mjs';

// An opt-in metadata discrimination evaluation, not native activation proof.
const cases = JSON.parse(await readFile(join(root, 'evaluations/routing.json'), 'utf8'));
const catalog = [];
for (const name of (await readdir(join(root, 'skills'))).sort()) {
  const content = await readFile(join(root, 'skills', name, 'SKILL.md'), 'utf8');
  const { description } = parse(content.match(/^---\n([\s\S]*?)\n---/)[1]);
  catalog.push({ name, description });
}
const output = join(root, '.cache/routing');
await mkdir(output, { recursive: true });
const scratch = await mkdtemp(join(tmpdir(), 'charm-routing-'));
try {
  const schemaPath = join(scratch, 'schema.json');
  await writeFile(schemaPath, JSON.stringify({ type: 'object', properties: { choices: { type: 'array', items: { type: 'object', properties: { id: { type: 'string' }, skill: { type: ['string', 'null'], enum: [null, ...catalog.map(c => c.name)] } }, required: ['id', 'skill'], additionalProperties: false } } }, required: ['choices'], additionalProperties: false }));
  const prompt = 'Do not call tools, read files, or invoke other agents. For each request, select exactly one primary skill from the supplied catalog, or null when none applies. Return only the requested JSON.\nCatalog:\n' + JSON.stringify(catalog) + '\nRequests:\n' + JSON.stringify(cases.map(({ id, prompt }) => ({ id, prompt })));
  await writeFile(join(output, 'prompt.txt'), prompt);
  const started = Date.now();
  const result = await run('codex', ['exec', '--ignore-user-config', '--ephemeral', '--skip-git-repo-check', '--sandbox', 'read-only', '--color', 'never', '--json', '--output-schema', schemaPath, '-o', join(output, 'answer.json'), '-C', scratch, '-'], { cwd: scratch, input: prompt });
  await writeFile(join(output, 'response.jsonl'), result.stdout);
  await writeFile(join(output, 'execution.json'), JSON.stringify({ ...result, stdout: undefined, elapsedMs: Date.now() - started, cli: (await run('codex', ['--version'])).stdout.trim() }, null, 2) + '\n');
  if (result.code !== 0) throw new Error('Routing provider failed; see .cache/routing/execution.json');
  const { choices } = JSON.parse(await readFile(join(output, 'answer.json'), 'utf8'));
  const actual = new Map(choices.map(c => [c.id, c.skill]));
  if (actual.size !== cases.length || choices.length !== cases.length) throw new Error('Missing or duplicate routing answers');
  const results = cases.map(c => ({ ...c, actual: actual.get(c.id), passed: actual.get(c.id) === c.expected }));
  await writeFile(join(output, 'results.json'), JSON.stringify(results, null, 2) + '\n');
  for (const c of results) console.log(`${c.passed ? 'PASS' : 'FAIL'} ${c.id}: ${c.actual}`);
  if (results.some(c => !c.passed)) process.exitCode = 1;
} finally { await rm(scratch, { recursive: true, force: true }); }
