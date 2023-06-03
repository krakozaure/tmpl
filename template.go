package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func templateRun(input string) error {
	outputString, err := templateExecute(input)
	if err != nil {
		return err
	}
	fmt.Print(outputString)
	return nil
}

func templateExecute(input string) (string, error) {
	var (
		err          error
		outputBytes  bytes.Buffer
		outputString string
	)

	inputBytes, err := readInput(input)
	if err != nil {
		return "", fmt.Errorf("unable to read input %v\n%v\n", input, err)
	}

	tmpl := template.New(input)
	tmpl.Funcs(getFuncMap())

	tmpl, err = tmpl.Parse(string(inputBytes))
	if err != nil {
		return "", fmt.Errorf("unable to parse input\n%v\n", err)
	}

	if len(ctx) == 0 {
		loadContext()
	}

	err = tmpl.Execute(&outputBytes, ctx)
	if err != nil {
		return "", fmt.Errorf("unable to render template\n%v\n", err)
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
		inputBytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		inputBytes, err = ioutil.ReadFile(input)
	}
	return inputBytes, err
}
