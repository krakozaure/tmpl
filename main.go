package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
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

	output, err := executeTemplateFile(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(output)
}
