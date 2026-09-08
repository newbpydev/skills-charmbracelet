import { spawn } from 'node:child_process';
import { readdir } from 'node:fs/promises';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';

export const root = fileURLToPath(new URL('../', import.meta.url));
export async function filesUnder(directory) {
  const files = [];
  for (const item of await readdir(directory, { withFileTypes: true })) {
    if (['node_modules', '.git', '.cache'].includes(item.name)) continue;
    const path = join(directory, item.name);
    if (item.isDirectory()) files.push(...await filesUnder(path));
    else files.push(path);
  }
  return files.sort();
}

export function run(command, args, { cwd = root, env = {}, timeout = 180_000, quiet = false, input } = {}) {
  return new Promise((resolveRun, reject) => {
    const child = spawn(command, args, { cwd, env: { ...process.env, ...env }, shell: false, stdio: [input === undefined ? 'ignore' : 'pipe', 'pipe', 'pipe'] });
    let stdout = '', stderr = '', timedOut = false;
    let forceKill;
    const timer = setTimeout(() => {
      timedOut = true;
      child.kill();
      forceKill = setTimeout(() => child.kill('SIGKILL'), 2000);
    }, timeout);
    if (input !== undefined) {
      child.stdin.on('error', () => {}); // The child may exit before consuming input.
      child.stdin.end(input);
    }
    child.stdout.on('data', data => { stdout += data; });
    child.stderr.on('data', data => { stderr += data; });
    child.on('error', error => { clearTimeout(timer); clearTimeout(forceKill); reject(error); });
    child.on('close', (code, signal) => {
      clearTimeout(timer);
      clearTimeout(forceKill);
      const result = { code: timedOut ? 124 : code, signal, timedOut, stdout, stderr };
      if (!quiet && result.code !== 0) process.stderr.write(stdout + stderr);
      resolveRun(result);
    });
  });
}
