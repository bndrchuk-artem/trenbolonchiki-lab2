package lab2

import (
	"io"
	"strconv"
	"strings"
	"log" // Import the "log" package

)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	data, err := io.ReadAll(ch.Input)
	if err != nil {
		log.Printf("Error reading from input: %v", err) // Log the error
		return err
	}

	expression := strings.TrimSpace(string(data)) // Trim leading/trailing whitespace

	if expression == "" {
		// Handle empty expression, potentially returning an error or a default value
		_, err = ch.Output.Write([]byte("0")) // Or return an error: return errors.New("empty expression")
		return err
	}

	result, err := EvaluatePostfix(expression)
	if err != nil {
		log.Printf("Error evaluating postfix expression: %v, expression: %s", err, expression) // Log the error and the expression
		return err
	}

	resultStr := strconv.Itoa(result)
	_, err = ch.Output.Write([]byte(resultStr))
	if err != nil {
		log.Printf("Error writing to output: %v, result: %s", err, resultStr) // Log the error and the result
		return err
	}

	return nil
}
