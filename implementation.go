package lab2

import (
	"errors"
	"strconv"
	"strings"
)

func EvaluatePostfix(input string) (int, error) {
	if input == "" {
		return 0, errors.New("empty input")
	}

	stack := []int{}
	tokens := strings.Fields(input)

	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, errors.New("Invalid postfix expression")
			}

			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			var result int
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, errors.New("Division by zero")
				}
				result = a / b
			case "^":
				result = 1
				for i := 0; i < b; i++ {
					result *= a
				}
			default:
				return 0, errors.New("Invalid operator")
			}

			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("Invalid postfix expression")
	}

	return stack[0], nil
}