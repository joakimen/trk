# trk

Tracks script usage.

## Description

`trk` is a simple script that tracks the usage of other scripts.

The script name (provided as the sole argument), as well as its invocation count, is recorded in `~/.local/share/trk/data.json`.

## Purpose

The purpose of this tool is to determine frequently used scripts, in order to determine which scripts and tools are no longer being used and may be safely removed.

## Usage


### Manually

`trk` supports executables in relative, absolute or $PATH commands.

Example (executed from `$HOME`):

```bash
trk pwd
trk ./bin/bar.sh
trk /bin/ls
trk /bin/ls
```

Result:

`~/.local/share/trk/data.json`:
```json
{
  "/Users/MyUser/bin/bar.sh": 1,
  "/bin/ls": 2,
  "pwd": 1,
}

```

### Automatically

#### zsh

In zsh, you can trigger `trk` automatically by adding the following line to your `.zshrc`:

```sh
preexec() { trk "$1" }
```

This way, `trk` will be triggered every time a command is executed. This should probably not be left on indefinitely, but for a period of time to determine your usage patterns and what scripts and tools may be safely removed.
