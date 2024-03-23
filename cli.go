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
	KeysList  stringsArray
	Strict    bool
)

func initFlags() {
	flag.Var(
		&FilesList,
		"f",
		"Load variables from one or more JSON/TOML/YAML files (format: file path)",
	)
	flag.Var(
		&KeysList,
		"k",
		"Load data from JSON/TOML/YAML files into keys (format: key=file path)",
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

	flag.Usage = func() {
		log.Printf(
			`USAGE: %s [OPTIONS] INPUT

INPUT is a template file or '-' for stdin

OPTIONS:
`,
			os.Args[0],
		)
		flag.PrintDefaults()
	}

	flag.Parse()
}
