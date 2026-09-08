import { test } from 'node:test';
import assert from 'node:assert/strict';
import { run } from '../scripts/lib.mjs';

test('a failed tool remains a failure with its diagnostic output', async () => {
  const result = await run(process.execPath, ['-e', "console.error('deliberate failure'); process.exit(7)"], { quiet: true });
  assert.equal(result.code, 7);
  assert.match(result.stderr, /deliberate failure/);
});
test('a deadline never becomes success when the tool exits cleanly after termination', async () => {
  const result = await run(process.execPath, ['-e', "process.on('SIGTERM', () => process.exit(0)); setInterval(() => {}, 100)"], { timeout: 500, quiet: true });
  assert.equal(result.code, 124);
  assert.equal(result.timedOut, true);
});
