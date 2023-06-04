package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	initFlags()
	if flag.NArg() == 0 {
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
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}
	fmt.Print(output)
}
