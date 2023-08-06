# tmpl

`tmpl` allows to apply variables from JSON/TOML/YAML files,
environment variables or CLI arguments to template files using Golang
`text/template` and functions from the Sprig project.

## Project status

This project is in a very early stage. Things might break !

## Usage

For simplicity, everything is output on `stdout`. To write the result in a
file, use shell redirection or the `-o` flag.

If a variable is not found and the strict mode is disabled (default disabled),
each missing value will be replaced with an empty string.

- Print help/usage message.

```bash
# $ tmpl
# $ tmpl -h
$ tmpl --help
USAGE: tmpl [OPTIONS] INPUT

INPUT is a template file or '-' for stdin

OPTIONS:
  -data string
        JSON data as a string or reference to a file ('@file') [DO NOT USE]
  -f value
        Load variables from one or more JSON/TOML/YAML files (format: file path)
  -o string
        Output file path, uses stdout if not set (format: file path)
  -s    Strict mode. Raise errors if variables are missing (default: false)
  -v value
        Use one or more variables from the command line (format: name=value)

The '-data' flag is for compatibility with https://github.com/benbjohnson/tmpl
and its use is not recommended.

When 'data' is used the '-f' and '-v' args are ignored, the INPUT file is
expected to have the '.tmpl' suffix and if the '-o' flag is not used the output
is written to the INPUT path removing the '.tmpl' suffix (if INPUT is '-' the
output goes to stdout).
```

- `stdin` and environment variables.

```bash
$ echo 'Editor = {{ env "EDITOR" }}, Shell = {{ env "SHELL" }}' | tmpl -
Editor = nvim, Shell = /bin/bash
```

- `stdin` and CLI variables.

```bash
$ tmpl -v foo=bar - <<< 'Hello, {{ .foo }} !'
Hello, bar !
```

- Sample from this repository

Output is not pasted here because of its length.

```bash
$ tmpl -f ./stores/data.yaml ./inputs/sample.txt.tmpl
```

## Configuration

No configuration needed. Everything is done on the command line.

## About the data flag

The `-data` flag has been included to be able the replace the `tmpl` binary on
the Debian distribution.

It was built from the https://github.com/benbjohnson/tmpl repository and the
idea is that it should not be used, as it is not really needed.
