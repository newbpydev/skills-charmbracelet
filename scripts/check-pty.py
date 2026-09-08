#!/usr/bin/env python3
"""Opt-in Linux PTY smoke tests; only the child PTY's terminal state is changed."""

import errno
import fcntl
import json
import os
from pathlib import Path
import select
import signal
import struct
import subprocess
import sys
import termios
import time

ROOT = Path(__file__).resolve().parent.parent
OUTPUT = ROOT / ".cache" / "pty"


def exercise(name, ready, quit_key):
    module = ROOT / "skills" / f"charmbracelet-{name}" / "assets" / "example"
    binary = OUTPUT / name
    subprocess.run(["go", "build", "-mod=readonly", "-o", str(binary), "."], cwd=module, check=True, timeout=120)
    master, slave = os.openpty()
    original = termios.tcgetattr(slave)
    transcript = bytearray()
    process = None

    def resize(width, height):
        fcntl.ioctl(slave, termios.TIOCSWINSZ, struct.pack("HHHH", height, width, 0, 0))

    def controlling_terminal():
        os.setsid()
        fcntl.ioctl(0, termios.TIOCSCTTY, 0)

    def drain(duration=0.1):
        deadline = time.monotonic() + duration
        while time.monotonic() < deadline:
            readable, _, _ = select.select([master], [], [], max(0, min(0.1, deadline - time.monotonic())))
            if readable:
                try:
                    data = os.read(master, 65536)
                except OSError as error:
                    if error.errno == errno.EIO:
                        return
                    raise
                if not data:
                    return
                transcript.extend(data)

    def wait_for(text, since=0):
        deadline = time.monotonic() + 8
        while text not in transcript[since:]:
            if process.poll() is not None or time.monotonic() >= deadline:
                raise AssertionError(f"{name}: did not render {text!r}; exit={process.poll()}")
            drain()

    try:
        resize(60, 12)
        process = subprocess.Popen(
            [str(binary)], cwd=module, stdin=slave, stdout=slave, stderr=slave,
            env={**os.environ, "TERM": "xterm-256color", "COLORTERM": "truecolor"},
            preexec_fn=controlling_terminal,
        )
        wait_for(ready)
        resize(2, 2)
        drain(0.2)
        before_resize = len(transcript)
        resize(80, 20)
        wait_for(ready, before_resize)
        if name == "bubbletea":
            os.write(master, b"tea\r")
            wait_for(b"Result: tea")
            os.write(master, b"q")
            drain(0.2)
            if process.poll() is not None:
                raise AssertionError("q quit while editing the focused input")
        os.write(master, quit_key)
        deadline = time.monotonic() + 5
        while process.poll() is None and time.monotonic() < deadline:
            drain()
        if process.poll() is None:
            raise AssertionError(f"{name}: quit timed out")
        drain()
        if process.returncode != 0:
            raise AssertionError(f"{name}: exit={process.returncode}")
        if termios.tcgetattr(slave) != original:
            raise AssertionError(f"{name}: terminal attributes were not restored")
        for sequence in (b"\x1b[?1049h", b"\x1b[?1049l"):
            if sequence not in transcript:
                raise AssertionError(f"{name}: missing alternate-screen transition {sequence!r}")
        return {"example": name, "exit": 0, "resize": ["60x12", "2x2", "80x20"], "input_and_quit": True, "termios_restored": True, "alternate_screen_restored": True}
    finally:
        if process is not None and process.poll() is None:
            os.killpg(process.pid, signal.SIGKILL)
            process.wait(timeout=5)
        (OUTPUT / f"{name}.terminal.bin").write_bytes(transcript)
        os.close(master)
        os.close(slave)


def main():
    if sys.platform != "linux":
        raise SystemExit("This PTY check is Linux-only; no terminal result was recorded.")
    OUTPUT.mkdir(parents=True, exist_ok=True)
    results = []
    for name, ready, quit_key in [("bubbletea", b"Enter: load", b"\x03"), ("terminal", "Hello 日本語".encode(), b"q")]:
        result = exercise(name, ready, quit_key)
        results.append(result)
        print(f"PASS {name}: input, resize, quit, termios and alternate-screen restoration")
    (OUTPUT / "results.json").write_text(json.dumps({"platform": sys.platform, "results": results}, indent=2) + "\n")


if __name__ == "__main__":
    main()
