# which

A utility designed to locate executable files in the current environment.

## Examples

**UNIX:**

```
$ which
Usage: which [filename ...]

$ which bash
/usr/bin/bash
```

**Windows:**

```
> which which
C:\Users\%USERPROFILE%\go\bin\which.exe

> which pwsh go node
C:\Program Files\PowerShell\7\pwsh.exe
C:\Program Files\Go\bin\go.exe
C:\Program Files\nodejs\node.exe
```

## Installation

```
go install github.com/chtozamm/cli-utilities/cmd/which@latest
```
