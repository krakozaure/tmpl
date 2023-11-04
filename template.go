package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"text/template"
)

func executeTemplateFile(input string) (string, error) {
	var (
		err          error
		outputBytes  bytes.Buffer
		outputString string
	)

	inputBytes, err := readInput(input)
	if err != nil {
		return "", err
	}

	tmpl := template.New(input)
	tmpl.Funcs(getFuncMap())

	tmpl, err = tmpl.Parse(string(inputBytes))
	if err != nil {
		return "", err
	}
	if Strict == true {
		tmpl.Option("missingkey=error")
	}

	if len(ctx) == 0 {
		loadContext()
	}

	err = tmpl.Execute(&outputBytes, ctx)
	if err != nil {
		return "", err
	}

	outputString = outputBytes.String()
	outputString = strings.ReplaceAll(outputString, "<no value>", "")
	return outputString, nil
}

func readInput(input string) ([]byte, error) {
	var (
		err        error
		inputBytes []byte
	)
	if input == "-" {
		inputBytes, err = io.ReadAll(os.Stdin)
	} else {
		inputBytes, err = os.ReadFile(input)
	}
	return inputBytes, err
}
