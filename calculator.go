// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Add takes two or more numbers and returns the result of adding them together.
func Add(numbers ...float64) float64 {
	sum := 0.0
	for _, n := range numbers {
		sum = sum + n
	}
	return sum
}

// Subtract takes two or numbers and returns the result of subtracting the first
// from the rest of the numbers.
func Subtract(a float64, numbers ...float64) float64 {
	value := a
	for _, n := range numbers {
		value = value - n
	}
	return value
}

// Multiple takes two or more numbers and returns the result of multiplying them
// together.
func Multiply(numbers ...float64) float64 {
	value := 1.0
	for _, n := range numbers {
		value = value * n
	}
	return value
}

// Divide takes atleast two numbers and returns the result of dividing the first
// number by the second and repeats division with subsequent numbers.
func Divide(a float64, numbers ...float64) (float64, error) {
	value := a
	for _, n := range numbers {
		if n == 0 {
			return 0, fmt.Errorf("bad input: %f, %v (division by zero is undefined)", a, numbers)
		}
		value = value / n
	}

	return value, nil
}

// Sqrt takes a number and return the square root of it.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input: %f (square root of a negative is undefined)", a)
	}
	return math.Sqrt(a), nil
}

// Evaluate takes a arithmetic expression string, and returns the evaluation of the
// arithmetic expression
func Evaluate(s string) (float64, error) {
	tokens := strings.Split(s, " ")
	if len(tokens) != 3 {
		return 0, fmt.Errorf("bad input: %q (unable to parse arithmetic expression)", s)
	}

	operand1, operator, operand2 := tokens[0], tokens[1], tokens[2]

	op1, err := strconv.ParseFloat(operand1, 64)
	if err != nil {
		return 0, fmt.Errorf("bad input: %q (first operand cannot be parsed as float)", operand1)
	}

	op2, err := strconv.ParseFloat(operand2, 64)
	if err != nil {
		return 0, fmt.Errorf("bad input: %q (first operand cannot be parsed as float)", operand2)
	}

	switch {
	case operator == "+":
		return Add(op1, op2), nil
	case operator == "-":
		return Subtract(op1, op2), nil
	case operator == "*":
		return Multiply(op1, op2), nil
	case operator == "/":
		val, err := Divide(op1, op2)
		if err != nil {
			return 0, err
		}
		return val, nil
	default:
		return 0, fmt.Errorf("bad operator input: %q (arithmetic operator is undefined)", operator)
	}
}
