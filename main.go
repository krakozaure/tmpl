package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const cdataExtension = ".tmpl"

func main() {
	var outputPath string

	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", os.Args[0]))

	initFlags()
	if flag.NArg() == 0 {
		log.SetPrefix("")
		flag.Usage()
		os.Exit(2)
	}

	input := flag.Arg(0)
	if input != "-" {
		inputAbs, err := filepath.Abs(input)
		if err == nil {
			input = inputAbs
		}
		inputsDir = filepath.Dir(inputAbs)
	}

	outputPath = OutPath
	if outputPath == "" && DataVar != "" && input != "-" {
		if !strings.HasSuffix(input, cdataExtension) {
			err := fmt.Errorf("path must have %s extension: %s", cdataExtension, input)
			log.Fatal(err)
		}
		outputPath = strings.TrimSuffix(input, cdataExtension)
	}

	output, err := executeTemplateFile(input)
	if err != nil {
		log.Fatal(err)
	}

	if outputPath == "" {
		fmt.Print(output)
	} else {
		fi, err := os.Stat(input)
		if err == nil {
			os.WriteFile(outputPath, []byte(output), fi.Mode())
		} else {
			log.Fatal(err)
		}
	}
}
