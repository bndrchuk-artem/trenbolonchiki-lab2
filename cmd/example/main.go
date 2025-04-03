package main

import (
	"flag"
	"fmt"
	"io"
	"github.com/bndrchuk-artem/trenbolonchiki-lab2"
	"os"
	"strings"
)

var (
	inputExpr = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "Input file with expression")
	outputFile = flag.String("o", "", "Output file for results")
)

func main() {
	flag.Parse()

	if *inputExpr != "" && *inputFile != "" {
		fmt.Fprintln(os.Stderr, "Error: use either -e or -f, not both")
		os.Exit(1)
	}

	var input io.Reader
	if *inputExpr != "" {
		input = strings.NewReader(*inputExpr)
	} else if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		fmt.Fprintln(os.Stderr, "Error: no input provided (use -e or -f)")
		os.Exit(1)
	}

	var output io.Writer
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	handler := &lab2.ComputeHandler{Input: input, Output: output}
	if err := handler.Compute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
