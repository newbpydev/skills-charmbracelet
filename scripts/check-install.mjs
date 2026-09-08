import { mkdtemp, mkdir, rm, readdir, readFile, realpath, lstat } from 'node:fs/promises';
import { join, relative, resolve, sep } from 'node:path';
import { tmpdir } from 'node:os';
import { filesUnder, root, run } from './lib.mjs';
import { validateSkill } from './validate-skills.mjs';

const targets = {
  'claude-code': '.claude/skills', codex: '.agents/skills', cursor: '.agents/skills',
  'gemini-cli': '.agents/skills', opencode: '.agents/skills', 'github-copilot': '.agents/skills',
  windsurf: '.windsurf/skills', cline: '.agents/skills', crush: '.crush/skills',
};
const names = (await readdir(join(root, 'skills'), { withFileTypes: true })).filter(e => e.isDirectory()).map(e => e.name).sort();
const cli = join(root, 'node_modules/skills/bin/cli.mjs');
const [source = root, ...extra] = process.argv.slice(2);
if (extra.length || source.startsWith('-')) throw new Error('Usage: node scripts/check-install.mjs [source]');
console.log(`Checking Skills CLI source: ${source}`);
const env = { DISABLE_TELEMETRY: '1', NO_COLOR: '1', FORCE_COLOR: '0', CI: '1' };
const scratch = await realpath(await mkdtemp(join(tmpdir(), 'charm-install-')));
let checks = 0;
async function install(args, cwd) {
  const result = await run(process.execPath, [cli, 'add', source, ...args], { cwd, env });
  if (result.code !== 0) throw new Error(`Skills CLI failed: ${result.stderr}`);
  return result;
}
async function verify(installed, name) {
  const errors = await validateSkill(installed);
  if (errors.length) throw new Error(`${name}: ${errors.join('; ')}`);
  const destination = await realpath(installed);
  if (!destination.startsWith(scratch + sep)) throw new Error(`Installation escaped temporary project: ${destination}`);
  const source = join(root, 'skills', name);
  for (const file of await filesUnder(source)) {
    const copied = resolve(installed, relative(source, file));
    if (!(await readFile(file)).equals(await readFile(copied))) throw new Error(`Resource changed or missing: ${copied}`);
  }
  checks++;
}
try {
  const listing = await install(['--list'], scratch);
  for (const name of names) if (!listing.stdout.includes(name)) throw new Error(`Not discovered: ${name}`);
  for (const copy of [false, true]) {
    const project = join(scratch, copy ? 'copy' : 'symlink');
    await mkdir(project);
    // CLI 1.5.24 skips non-universal project symlinks (except Claude) when
    // their harness directory does not already exist. Copy mode creates it.
    if (!copy) for (const path of new Set(Object.values(targets))) {
      await mkdir(join(project, path.split('/')[0]), { recursive: true });
    }
    await install(['--skill', '*', '--agent', ...Object.keys(targets), '--yes', ...(copy ? ['--copy'] : [])], project);
    for (const [target, path] of Object.entries(targets)) {
      let links = 0;
      for (const name of names) {
        const installed = join(project, path, name);
        await verify(installed, name);
        if ((await lstat(installed)).isSymbolicLink()) links++;
      }
      console.log(`PASS ${target} ${copy ? 'copy' : 'default mode, existing harness directory'} (${names.length} skills, ${links} links)`);
    }
  }
  for (const name of names) {
    const project = join(scratch, name); await mkdir(project);
    await install(['--skill', name, '--agent', 'codex', '--copy', '--yes'], project);
    await verify(join(project, '.agents/skills', name), name);
    const installedNames = await readdir(join(project, '.agents/skills'));
    if (installedNames.length !== 1 || installedNames[0] !== name) throw new Error(`Standalone install included other skills: ${name}`);
  }
  console.log(`PASS standalone installs; ${checks} installed skill/resource checks. Real global directories were not used.`);
} finally { await rm(scratch, { recursive: true, force: true }); }
