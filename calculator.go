// Package calculator provides a library for simple calculations in Go.
package calculator

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the first
// from the second.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiple takes two numbers and returns the result of multiplying them together.
func Multiply(a, b float64) float64 {
	return a * b
}
