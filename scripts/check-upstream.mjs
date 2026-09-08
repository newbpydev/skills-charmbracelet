import { readFile } from 'node:fs/promises';
import { join } from 'node:path';
import { root } from './lib.mjs';

const manifest = JSON.parse(await readFile(join(root, 'upstream-sources.json'), 'utf8'));
const failures = [], drift = [], checked = new Set();
async function get(url) {
  const response = await fetch(url, { signal: AbortSignal.timeout(20_000), headers: { 'User-Agent': 'skills-charmbracelet-validation' } });
  if (!response.ok) throw new Error(`HTTP ${response.status} ${url}`);
  return response;
}
for (const source of manifest.modules) {
  try {
    const latest = await (await get(`https://proxy.golang.org/${source.module}/@latest`)).json();
    if (latest.Version !== source.version) drift.push(`${source.module}: pinned ${source.version}, upstream ${latest.Version}`);
    for (const path of [...source.sources, source.licensePath]) {
      const url = `https://raw.githubusercontent.com/${source.repository}/${source.commit}/${path}`;
      if (checked.has(url)) continue;
      checked.add(url);
      await (await get(url)).arrayBuffer();
    }
  } catch (error) { failures.push(`${source.id}: ${error.message}`); }
}
for (const item of drift) console.log(`REVIEW ${item}`);
for (const failure of failures) console.error(`UNVERIFIED ${failure}`);
console.log(`Checked ${manifest.modules.length} module baselines and ${checked.size} pinned source URLs; ${drift.length} release changes, ${failures.length} unavailable sources.`);
if (failures.length) process.exitCode = 1;
