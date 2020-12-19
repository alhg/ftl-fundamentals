package calculator_test

import (
	"calculator"
	"strconv"
	"strings"
	"testing"
)

type testCase struct {
	name        string
	inputs      []float64
	want        float64
	errExpected bool
}

func getFormatVariadicInputStr(inputs []float64) string {
	stringArr := []string{}
	for _, i := range inputs {
		stringArr = append(stringArr, strconv.FormatFloat(i, 'f', -1, 64))
	}
	return strings.Join(stringArr, ", ")
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "Two positive numbers that sum to a positive", inputs: []float64{2, 2}, want: 4},
		{name: "A negative & postive number that sum to positive", inputs: []float64{-1, 1}, want: 0},
		{name: "A postive and negative that sum to negative", inputs: []float64{4, -5}, want: -1},
		{name: "Two negative numbers that sum to a negative", inputs: []float64{-1, -1}, want: -2},
		{name: "Ten positive and negative numbers that sum to a positive", inputs: []float64{-2, 5, 2, -4, 3}, want: 4},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.inputs...)
		if tc.want != got {
			t.Errorf("%s\nAdd(%s): want %f, got %f", tc.name, getFormatVariadicInputStr(tc.inputs), tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "Two positive numbers that subtract to a positive", inputs: []float64{4, 2}, want: 2},
		{name: "Two positive numbers that subtract to a negative", inputs: []float64{3, 4}, want: -1},
		{name: "Two negative numbers that subtract to a positive", inputs: []float64{-1, -2}, want: 1},
		{name: "Two negative numbers that subtract to a negative", inputs: []float64{-3, -2}, want: -1},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.inputs[0], tc.inputs[1:]...)
		if tc.want != got {
			t.Errorf("%s\nSubtract(%s): want %f, got %f", tc.name, getFormatVariadicInputStr(tc.inputs), tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "Two postives numbers that multiply to positive", inputs: []float64{3, 3}, want: 9},
		{name: "Two negative numbers the multiply to positive", inputs: []float64{-2, -3}, want: 6},
		{name: "A negative and postive (in that order) that multply to negative", inputs: []float64{-4, 2}, want: -8},
		{name: "A positive and negative (in that order) that multiple to negative", inputs: []float64{5, -1}, want: -5},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.inputs...)
		if tc.want != got {
			t.Errorf("%s\nMultiply(%s): want %f, got %f", tc.name, getFormatVariadicInputStr(tc.inputs), tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "Two postive numbers divided returns a postive", inputs: []float64{4, 2}, want: 2, errExpected: false},
		{name: "Two negative numbers divided returns a positive", inputs: []float64{-4, -2}, want: 2, errExpected: false},
		{name: "A positive and negative numbers divided return a negative", inputs: []float64{5, -2}, want: -2.5, errExpected: false},
		{name: "A number divided by zero that returns an error", inputs: []float64{2, 0}, want: 0, errExpected: true},
	}
	for _, tc := range testCases {

		got, err := calculator.Divide(tc.inputs[0], tc.inputs[1:]...)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("%s\nDivide(%s): unexpected error status: %v", tc.name, getFormatVariadicInputStr(tc.inputs), errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("%s\nDivide(%s): want %f, got %f", tc.name, getFormatVariadicInputStr(tc.inputs), tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type singleInputTestCase struct {
		name        string
		input       float64
		want        float64
		errExpected bool
	}
	testCases := []singleInputTestCase{
		{name: "A positive number returns its square root", input: 4, want: 2, errExpected: false},
		{name: "The square root of 0 will return 0", input: 0, want: 0, errExpected: false},
		{name: "A negative number returns an error", input: -2, want: 0, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.input)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("%s\nSqrt(%f): unexpected error status: %v", tc.name, tc.input, errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("%s\nSqrt(%f): want %f, got %f", tc.name, tc.input, tc.want, got)
		}
	}
}
