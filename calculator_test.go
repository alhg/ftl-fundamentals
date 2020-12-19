package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	name string
	a, b float64
	want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "Two positive numbers that sum to a positive", a: 2, b: 2, want: 4},
		{name: "A negative & postive number that sum to positive", a: -1, b: 1, want: 0},
		{name: "A postive and negative that sum to negative", a: 4, b: -5, want: -1},
		{name: "Two negative numbers that sum to a negative", a: -1, b: -1, want: -2},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s\nAdd(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "Two positive numbers that subtract to a positive", a: 4, b: 2, want: 2},
		{name: "Two positive numbers that subtract to a negative", a: 3, b: 4, want: -1},
		{name: "Two negative numbers that subtract to a positive", a: -1, b: -2, want: 1},
		{name: "Two negative numbers that subtract to a negative", a: -3, b: -2, want: -1},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s\nSubtract(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "Two postives numbers that multiply to positive", a: 3, b: 3, want: 9},
		{name: "Two negative numbers the multiply to positive", a: -2, b: -3, want: 6},
		{name: "A negative and postive (in that order) that multply to negative", a: -4, b: 2, want: -8},
		{name: "A positive and negative (in that order) that multiple to negative", a: 5, b: -1, want: -5},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s\nMultiply(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}
