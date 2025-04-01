package main

import (
	"flag"
	"fmt"
	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	// TODO: Add other flags support for input and output configuration.
)

func main() {
	flag.Parse()

	
	      handler := &lab2.ComputeHandler{
	          Input: {construct io.Reader according the command line parameters},
	          Output: {construct io.Writer according the command line parameters},
	      }
	      err := handler.Compute()

	res, _ := lab2.PrefixToPostfix("+ 2 2")
	fmt.Println(res)
}
