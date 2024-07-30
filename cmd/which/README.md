# which

A tool to locate a command in the current environment.
It does this by searching the `PATH` for executable files matching the names of the arguments.

## Example

**UNIX:**

```sh
$ which
Usage: which [filename ...]

$ which bash
/usr/bin/bash
```

**Windows:**

```sh
PS which which
C:\Users\$env:USERPROFILE\go\bin\which.exe

PS which pwsh go node
C:\Program Files\PowerShell\7\pwsh.exe
C:\Program Files\Go\bin\go.exe
C:\Program Files\nodejs\node.exe
```
