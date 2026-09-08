import { test } from 'node:test';
import assert from 'node:assert/strict';
import { mkdtemp, mkdir, writeFile, rm } from 'node:fs/promises';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { validateSkill, validateCatalog } from '../scripts/validate-skills.mjs';

async function fixture(t, body, frontmatter = 'name: example\ndescription: Build an example.') {
  const base = await mkdtemp(join(tmpdir(), 'charm-validation-'));
  t.after(() => rm(base, { recursive: true, force: true }));
  const dir = join(base, 'example');
  await mkdir(dir);
  await writeFile(join(dir, 'SKILL.md'), `---\n${frontmatter}\n---\n${body}\n`);
  return dir;
}

test('a standalone skill with a real reference is valid', async t => {
  const dir = await fixture(t, '[Details](references/details.md)');
  await mkdir(join(dir, 'references'));
  await writeFile(join(dir, 'references/details.md'), '# Details\n');
  assert.deepEqual(await validateSkill(dir), []);
});
test('missing resources fail after a standalone install', async t => {
  const dir = await fixture(t, '[Details](references/missing.md)');
  assert.match((await validateSkill(dir)).join('\n'), /Broken local link/);
});
test('sibling dependencies violate standalone packaging', async t => {
  const dir = await fixture(t, '[Shared](../shared.md)');
  assert.match((await validateSkill(dir)).join('\n'), /escapes skill/);
});
test('name mismatches and empty descriptions are rejected', async t => {
  const dir = await fixture(t, 'Content', 'name: other\ndescription: ""');
  assert.equal((await validateSkill(dir)).length, 2);
});
test('undiscoverable conditional references are rejected', async t => {
  const dir = await fixture(t, 'Content');
  await mkdir(join(dir, 'references'));
  await writeFile(join(dir, 'references/hidden.md'), 'Hidden');
  assert.match((await validateSkill(dir)).join('\n'), /not linked directly/);
});
test('reference-style links are resolved', async t => {
  const dir = await fixture(t, '[Guide][guide]\n\n[guide]: absent.md');
  assert.match((await validateSkill(dir)).join('\n'), /Broken local link/);
});
test('malformed URL escapes fail as validation errors', async t => {
  const dir = await fixture(t, '[Guide](references/%zz.md)');
  assert.match((await validateSkill(dir)).join('\n'), /Malformed link/);
});

test('catalog typos cannot silently omit a real skill', () => {
  const config = { groupings: [{ title: 'Core', skills: ['bubbleteaa'] }] };
  assert.deepEqual(validateCatalog(config, ['bubbletea']), [
    'Unknown catalog skill: bubbleteaa', 'Ungrouped catalog skill: bubbletea',
  ]);
});
test('catalog membership is complete and unambiguous', () => {
  const core = { title: 'Core', skills: ['bubbletea'] };
  assert.deepEqual(validateCatalog({ groupings: [core] }, ['bubbletea']), []);
  assert.deepEqual(validateCatalog({ groupings: [core, core] }, ['bubbletea']), ['Duplicate catalog skill: bubbletea']);
});
