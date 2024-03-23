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
  -f value
        Load variables from one or more JSON/TOML/YAML files (format: file path)
  -k value
        Load data from JSON/TOML/YAML files into keys (format: key=file path)
  -s    Strict mode. Raise errors if variables are missing (default: false)
  -v value
        Use one or more variables from the command line (format: name=value)
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
$ tmpl -f ./stores/data.yaml -k Key=./stores/list.yaml ./inputs/sample.txt.tmpl
```

## Configuration

No configuration needed. Everything is done on the command line.
