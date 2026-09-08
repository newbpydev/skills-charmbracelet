import { mkdtemp, mkdir, readFile, readdir, writeFile, rm } from 'node:fs/promises';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { root, run } from './lib.mjs';

// This check invokes an authenticated local agent. It is opt-in, outside npm run check.
const scratch = await mkdtemp(join(tmpdir(), 'charm-native-'));
const output = join(root, '.cache/native');
await mkdir(output, { recursive: true });
const names = (await readdir(join(root, 'skills'))).filter(n => n.startsWith('charmbracelet-')).sort();
try {
  const install = await run(process.execPath, [join(root, 'node_modules/skills/bin/cli.mjs'), 'add', root, '--skill', '*', '--agent', 'codex', '--copy', '--yes'], { cwd: scratch, env: { DISABLE_TELEMETRY: '1' } });
  if (install.code !== 0) throw new Error('Temporary skill installation failed');
  const initialized = await run('git', ['init', '--quiet'], { cwd: scratch });
  if (initialized.code !== 0) throw new Error('Temporary Git initialization failed');
  for (const [id, prompt] of [
    ['discovery', 'Without reading files or calling tools, list only the available skill names that begin with charmbracelet-, one per line. Use the skills available in this session, not guesses.'],
    ['navigation', 'Use $charmbracelet-huh. Read its entrypoint and relevant local references. Explain the pinned Huh Form.Update return type and accessible-mode value/error handling, with local file citations. Do not modify files, access the network, or invoke other agents.'],
  ]) {
    const answer = join(output, `${id}.txt`);
    const result = await run('codex', ['exec', '--ignore-user-config', '--ephemeral', '--skip-git-repo-check', '--sandbox', 'read-only', '--color', 'never', '--json', '-C', scratch, '-o', answer, prompt], { cwd: scratch, timeout: 180_000 });
    await writeFile(join(output, `${id}.jsonl`), result.stdout);
    await writeFile(join(output, `${id}-execution.json`), JSON.stringify({ code: result.code, signal: result.signal, stderr: result.stderr, cli: (await run('codex', ['--version'])).stdout.trim() }, null, 2) + '\n');
    if (result.code !== 0) throw new Error(`Native ${id} check failed`);
    const text = await readFile(answer, 'utf8');
    if (id === 'discovery') {
      const found = [...new Set(text.match(/charmbracelet-[a-z-]+/g) ?? [])].sort();
      if (JSON.stringify(found) !== JSON.stringify(names)) throw new Error(`Native discovery mismatch: ${found.join(', ')}`);
    }
    console.log(`Completed Codex native ${id}; evidence: ${answer}`);
  }
} finally { await rm(scratch, { recursive: true, force: true }); }
