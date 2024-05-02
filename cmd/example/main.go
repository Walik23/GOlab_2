package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
    "GOlab_2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	fileExpression  = flag.String("f", "", "File with expression")
	outputFile      = flag.String("o", "", "Output file")
)

func main() {
	flag.Parse()

	var input io.Reader
	var output io.Writer

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else if *fileExpression != "" {
		file, err := os.Open(*fileExpression)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening input file:", err)
			return
		}
		defer file.Close()
		input = file
	} else {
		fmt.Fprintln(os.Stderr, "No input provided")
		return
	}

	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating output file:", err)
			return
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	handler := lab2.ComputeHandler{
		Input:      input,
		Output:     output,
		Calculator: &lab2.DefaultInfixInterpreter{},
	}

	if err := handler.Compute(); err != nil {
		log.Fatalln("main: error occurred for Compute(), error: ", err)
	}
}

