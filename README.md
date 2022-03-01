# tmpl

`tmpl` allows to apply variables from JSON/TOML/YAML files,
environment variables or CLI arguments to template files using Golang
`text/template` and functions from the Sprig project.

## Project status

This project is in a very early stage. Things might break !

## Usage

For simplicity, everything is output on `stdout`. To write the result in a
file, use shell redirection.

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
  -e    load variables from environment (default true)
  -f value
        load variables from JSON/TOML/YAML files (format: file path)
  -s    exit on any error during template processing (default false)
  -v value
        use one or more variables from the command line (format: name=value)
```

- `stdin` and environment variables.

```bash
$ echo "Editor = {{ .Env.EDITOR }}, Shell = {{ .Env.SHELL }}" | tmpl -
Editor = nvim, Shell = /bin/bash
```

- `stdin` and CLI variables.

```bash
$ echo "Hello, {{ .foo }} !" | tmpl -v foo=bar -
Hello, bar !
```

- Sample from this repository

Output is not pasted here because of its length.

```bash
$ tmpl -f ./stores/data.yaml ./inputs/sample.txt.tmpl
```

## Configuration

No configuration needed. Everything is done on the command line.
