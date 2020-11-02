package lab2

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// EvaluatePostfix is function for calculation postfix notation
func EvaluatePostfix(input string) (float64, error) {

	evaluateOperators := map[string]func(float64, float64) float64{
		"+": func(arg1, arg2 float64) float64 {
			return arg1 + arg2
		},
		"-": func(arg1, arg2 float64) float64 {
			return arg1 - arg2
		},
		"*": func(arg1, arg2 float64) float64 {
			return arg1 * arg2
		},
		"/": func(arg1, arg2 float64) float64 {
			return arg1 / arg2
		},
		"^": func(arg1, arg2 float64) float64 {
			return math.Pow(arg1, arg2)
		},
	}

	if input == "" {
		return 0, errors.New("Input can not be empty")
	}

	stringArray := strings.Fields(input)
	var stack []float64

	for _, token := range stringArray {
		if operators, consist := evaluateOperators[token]; consist {
			if len(stack) < 2 {
				return 0, errors.New("Missing operand")
			}
			arg1, arg2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, operators(arg1, arg2))
		} else {
			val, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, val)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("Stack corrupted")
	}

	return stack[len(stack)-1], nil
}
