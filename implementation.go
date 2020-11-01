package lab2

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func EvaluatePostfix(input string) (int, error) {

	evaluateOperators := map[string]func(int, int) int{
		"+": func(x, y int) int {
			return x + y
		},
		"-": func(x, y int) int {
			return x - y
		},
		"*": func(x, y int) int {
			return x * y
		},
		"/": func(x, y int) int {
			return x / y
		},
		"^": func(x, y int) int {
			return int(math.Pow(float64(x), float64(y)))
		},
	}

	if input == "" {
		return 0, errors.New("Input can not be empty")
	}

	stringArray := strings.Fields(input)
	var stack []int

	for _, token := range stringArray {
		if operators, consist := evaluateOperators[token]; consist {
			if len(stack) < 2 {
				return 0, errors.New("Missing operand")
			}
			arg1, arg2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, operators(arg1, arg2))
		} else {
			val, err := strconv.Atoi(token)
			if err != nil {
				return 0, err
			}
			stack = append(stack, val)
		}
	}

	if len(stack) != 1 {
		return -1, errors.New("Stack corrupted")
	}

	return stack[len(stack)-1], nil
}
