import { readFile } from 'node:fs/promises';
import { dirname, join, relative } from 'node:path';
import { filesUnder, root, run } from './lib.mjs';

const modules = (await filesUnder(join(root, 'skills'))).filter(p => p.endsWith('/go.mod') || p.endsWith('\\go.mod'));
let failures = 0;
const race = process.argv.includes('--race');
for (const mod of modules) {
  const cwd = dirname(mod);
  const before = await readFile(mod, 'utf8');
  const sumBefore = await readFile(join(cwd, 'go.sum'), 'utf8');
  const sources = (await filesUnder(cwd)).filter(p => p.endsWith('.go'));
  const formatted = await run('gofmt', ['-l', ...sources]);
  if (formatted.code !== 0 || formatted.stdout.trim()) { console.error('Formatting:', formatted.stdout); failures++; }
  const test = await run('go', ['test', '-mod=readonly', '-timeout=45s', ...(race ? ['-race'] : []), './...'], { cwd });
  const vet = await run('go', ['vet', '-mod=readonly', './...'], { cwd });
  if (test.code !== 0 || vet.code !== 0) failures++;
  if (before !== await readFile(mod, 'utf8') || sumBefore !== await readFile(join(cwd, 'go.sum'), 'utf8')) {
    console.error('Dependency files changed during verification:', relative(root, cwd)); failures++;
  }
  console.log(`${test.code === 0 && vet.code === 0 ? 'PASS' : 'FAIL'} ${relative(root, cwd)}${race ? ' (race)' : ''}`);
}
console.log(`Checked ${modules.length} independent Go modules; ${failures} failures.`);
if (failures || modules.length === 0) process.exitCode = 1;
