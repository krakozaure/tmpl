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
	UseEnv    bool
	Strict    bool
)

func initFlags() {

	flag.BoolVar(
		&UseEnv,
		"e",
		true,
		"load variables from environment",
	)
	flag.Var(
		&FilesList,
		"f",
		"load variables from JSON/TOML/YAML files (format: file path)",
	)
	flag.Var(
		&VarsList,
		"v",
		"use one or more variables from the command line (format: name=value)",
	)

	flag.BoolVar(
		&Strict,
		"s",
		false,
		"exit on any error during template processing (default false)",
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
