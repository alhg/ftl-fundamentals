// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the first
// from the second.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiple takes two numbers and returns the result of multiplying them
// together.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing the first
// number by the second number.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("bad input: %f, %f (division by zero is undefined)", a, b)
	}
	return a / b, nil
}

// Sqrt takes a number and return the square root of it.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input: %f (square root of a negative is undefined)", a)
	}
	return math.Sqrt(a), nil
}
