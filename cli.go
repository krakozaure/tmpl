package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Allow multiple string values for one flag
type stringsArray []string

func (s *stringsArray) String() string {
	return fmt.Sprint(*s)
}

func (s *stringsArray) Set(value string) error {
	*s = append(*s, value)
	return nil
}

// Flags
var (
	VarsList  stringsArray
	FilesList stringsArray
	Strict    bool
	DataVar   string
	OutPath   string
)

func initFlags() {
	flag.Var(
		&FilesList,
		"f",
		"Load variables from one or more JSON/TOML/YAML files (format: file path)",
	)
	flag.Var(
		&VarsList,
		"v",
		"Use one or more variables from the command line (format: name=value)",
	)

	flag.BoolVar(
		&Strict,
		"s",
		false,
		"Strict mode. Raise errors if variables are missing (default: false)",
	)

	flag.StringVar(
		&DataVar,
		"data",
		"",
		"JSON data as a string or reference to a file ('@file') [DO NOT USE]",
	)

	flag.StringVar(
		&OutPath,
		"o",
		"",
		"Output file path, uses stdout if not set (format: file path)",
	)

	flag.Usage = func() {
		log.Printf(
			`USAGE: %s [OPTIONS] INPUT

INPUT is a template file or '-' for stdin

OPTIONS:
`,
			os.Args[0],
		)
		flag.PrintDefaults()
		log.Printf(`
The '-data' flag is for compatibility with https://github.com/benbjohnson/tmpl
and its use is not recommended.

When 'data' is used the '-f' and '-v' args are ignored, the INPUT file is
expected to have the '.tmpl' suffix and if the '-o' flag is not used the output
is written to the INPUT path removing the '.tmpl' suffix (if INPUT is '-' the
output goes to stdout).
`,
		)
	}

	flag.Parse()
}
