import { readFile, readdir, lstat, realpath } from 'node:fs/promises';
import { basename, dirname, isAbsolute, join, relative, resolve, sep } from 'node:path';
import { pathToFileURL } from 'node:url';
import { parse } from 'yaml';
import { fromMarkdown } from 'mdast-util-from-markdown';
import { filesUnder, root } from './lib.mjs';

function inside(base, path) {
  const rel = relative(base, path);
  return rel !== '..' && !rel.startsWith(`..${sep}`) && !isAbsolute(rel);
}
function links(markdown) {
  const tree = fromMarkdown(markdown);
  const definitions = new Map(), found = [], references = [];
  function walk(node) {
    if (node.type === 'definition') definitions.set(node.identifier.toLowerCase(), node.url);
    if (['link', 'image'].includes(node.type)) found.push(node.url);
    if (['linkReference', 'imageReference'].includes(node.type)) references.push(node.identifier.toLowerCase());
    node.children?.forEach(walk);
  }
  walk(tree);
  for (const ref of references) {
    if (!definitions.has(ref)) throw new Error(`Undefined Markdown reference: ${ref}`);
    found.push(definitions.get(ref));
  }
  return found;
}

export async function validateSkill(directory) {
  const errors = [];
  directory = await realpath(directory);
  const skillPath = join(directory, 'SKILL.md');
  let content;
  try { content = await readFile(skillPath, 'utf8'); }
  catch { return ['Missing SKILL.md']; }
  const match = content.match(/^---\r?\n([\s\S]*?)\r?\n---(?:\r?\n|$)/);
  if (!match) return ['Missing YAML frontmatter'];
  let metadata;
  try { metadata = parse(match[1], { uniqueKeys: true, maxAliasCount: 20 }); }
  catch (e) { return [`Invalid YAML: ${e.message}`]; }
  if (!metadata || typeof metadata !== 'object' || Array.isArray(metadata)) return ['Frontmatter must be a mapping'];
  if (typeof metadata.name !== 'string' || !/^[a-z0-9]+(?:-[a-z0-9]+)*$/.test(metadata.name) || metadata.name.length > 64 || metadata.name !== basename(directory)) errors.push('Name must match the directory and Agent Skills naming rules');
  if (typeof metadata.description !== 'string' || !metadata.description.trim() || metadata.description.length > 1024 || /[<>]/.test(metadata.description)) errors.push('Description must contain 1–1024 characters without XML tags');
  if (metadata.metadata && (Array.isArray(metadata.metadata) || typeof metadata.metadata !== 'object' || Object.values(metadata.metadata).some(v => typeof v !== 'string'))) errors.push('Metadata values must be strings');
  const body = content.slice(match[0].length);
  if (body.split('\n').length >= 500) errors.push('SKILL.md body must remain below 500 lines');
  if (!body.trim()) errors.push('Skill body is empty');
  if (/\[TODO[:\]]|\[INSERT\b/.test(content)) errors.push('Unfinished scaffold placeholder');
  const files = await filesUnder(directory);
  let directLinks = [];
  for (const file of files) {
    if ((await lstat(file)).isSymbolicLink()) errors.push(`Distributable file is a symlink: ${relative(directory, file)}`);
    if (!file.endsWith('.md')) continue;
    let urls;
    try { urls = links(file === skillPath ? body : await readFile(file, 'utf8')); }
    catch (e) { errors.push(`${relative(directory, file)}: ${e.message}`); continue; }
    for (const url of urls) {
      if (/^(?:https?:|mailto:|#)/i.test(url)) continue;
      let clean;
      try { clean = decodeURIComponent(url.split('#')[0]); }
      catch { errors.push(`Malformed link in ${relative(directory, file)}: ${url}`); continue; }
      const target = resolve(dirname(file), clean);
      if (!inside(directory, target)) { errors.push(`Link escapes skill: ${url}`); continue; }
      try {
        if (!inside(directory, await realpath(target))) errors.push(`Link target escapes skill: ${url}`);
      } catch { errors.push(`Broken local link in ${relative(directory, file)}: ${url}`); }
      if (file === skillPath) directLinks.push(target);
    }
  }
  for (const reference of files.filter(f => f.startsWith(join(directory, 'references') + sep) && f.endsWith('.md'))) {
    if (!directLinks.includes(reference)) errors.push(`Reference not linked directly from SKILL.md: ${relative(directory, reference)}`);
  }
  return errors;
}

export function validateCatalog(config, names) {
  const failures = [], seen = new Set();
  if (!Array.isArray(config?.groupings) || !config.groupings.length) return ['skills.sh.json needs at least one group'];
  for (const group of config.groupings) {
    if (!group || typeof group.title !== 'string' || !group.title.trim() || !Array.isArray(group.skills) || !group.skills.length) {
      failures.push('skills.sh.json group needs a title and skills');
      continue;
    }
    for (const name of group.skills) {
      if (!names.includes(name)) failures.push(`Unknown catalog skill: ${name}`);
      if (seen.has(name)) failures.push(`Duplicate catalog skill: ${name}`);
      seen.add(name);
    }
  }
  for (const name of names) if (!seen.has(name)) failures.push(`Ungrouped catalog skill: ${name}`);
  return failures;
}

export async function validateCollection(directory = root) {
  const skillRoot = join(directory, 'skills');
  const dirs = (await readdir(skillRoot, { withFileTypes: true })).filter(d => d.isDirectory()).map(d => d.name).sort();
  const failures = [];
  for (const name of dirs) for (const error of await validateSkill(join(skillRoot, name))) failures.push(`${name}: ${error}`);
  const catalog = JSON.parse(await readFile(join(directory, 'skills.sh.json'), 'utf8'));
  failures.push(...validateCatalog(catalog, dirs));
  const manifest = JSON.parse(await readFile(join(directory, 'upstream-sources.json'), 'utf8'));
  for (const source of manifest.modules) {
    if (!/^[a-f0-9]{40}$/.test(source.commit) || !source.version || !source.module || !source.license) failures.push(`Incomplete source pin: ${source.id}`);
    for (const consumer of source.consumers) {
      try { await readFile(join(directory, consumer)); } catch { failures.push(`Missing source consumer: ${consumer}`); }
    }
  }
  return { skills: dirs, failures };
}

if (process.argv[1] && import.meta.url === pathToFileURL(resolve(process.argv[1])).href) {
  const result = await validateCollection();
  for (const failure of result.failures) console.error(failure);
  console.log(`Validated ${result.skills.length} skills; ${result.failures.length} errors.`);
  if (result.failures.length) process.exitCode = 1;
}
