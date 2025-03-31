package lab2

import (
	"io"
	"strconv"
	"strings"
	"log"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	data, err := io.ReadAll(ch.Input)
	if err != nil {
		log.Printf("Error reading from input: %v", err)
		return err
	}

	expression := strings.TrimSpace(string(data))

	if expression == "" {
		_, err = ch.Output.Write([]byte("0"))
		return err
	}

	result, err := EvaluatePostfix(expression)
	if err != nil {
		log.Printf("Error evaluating postfix expression: %v, expression: %s", err, expression)
		return err
	}

	resultStr := strconv.Itoa(result)
	_, err = ch.Output.Write([]byte(resultStr))
	if err != nil {
		log.Printf("Error writing to output: %v, result: %s", err, resultStr) 
		return err
	}

	return nil
}
