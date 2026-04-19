package cmd

import (
	"fmt"
	"math"
	"strconv"
)

// parsePair parses two strings as float64, labeling errors with position
// so the user knows which argument was bad.
func parsePair(a, b string) (float64, float64, error) {
	x, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("first value %q is not a number", a)
	}
	y, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("second value %q is not a number", b)
	}
	return x, y, nil
}

// parseOne parses a single float argument.
func parseOne(a string) (float64, error) {
	x, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, fmt.Errorf("value %q is not a number", a)
	}
	return x, nil
}

// formatResult trims trailing zeros so 5.0 prints as "5" but 3.14159 stays
// intact. Whole-valued floats within the exact-integer range of float64 are
// rendered as plain integers (e.g. 10! -> "3628800" rather than "3.6288e+06").
func formatResult(v float64) string {
	if math.IsNaN(v) {
		return "NaN"
	}
	if math.IsInf(v, 1) {
		return "+Inf"
	}
	if math.IsInf(v, -1) {
		return "-Inf"
	}
	// 2^53 is the boundary above which consecutive integers aren't all
	// representable in float64; beyond that, 'f' formatting would lie.
	if v == math.Trunc(v) && math.Abs(v) < 1<<53 {
		return strconv.FormatFloat(v, 'f', 0, 64)
	}
	return strconv.FormatFloat(v, 'g', -1, 64)
}

// --- binary operations ---

func Add(a, b string) (string, error) {
	x, y, err := parsePair(a, b)
	if err != nil {
		return "", err
	}
	return formatResult(x + y), nil
}

func Subtract(a, b string) (string, error) {
	x, y, err := parsePair(a, b)
	if err != nil {
		return "", err
	}
	return formatResult(x - y), nil
}

func Multiply(a, b string) (string, error) {
	x, y, err := parsePair(a, b)
	if err != nil {
		return "", err
	}
	return formatResult(x * y), nil
}

func Divide(a, b string) (string, error) {
	x, y, err := parsePair(a, b)
	if err != nil {
		return "", err
	}
	if y == 0 {
		return "", fmt.Errorf("cannot divide by zero")
	}
	return formatResult(x / y), nil
}

func Power(base, exp string) (string, error) {
	x, y, err := parsePair(base, exp)
	if err != nil {
		return "", err
	}
	return formatResult(math.Pow(x, y)), nil
}

func Mod(a, b string) (string, error) {
	x, y, err := parsePair(a, b)
	if err != nil {
		return "", err
	}
	if y == 0 {
		return "", fmt.Errorf("cannot take modulo by zero")
	}
	return formatResult(math.Mod(x, y)), nil
}

// --- unary operations ---

func Sqrt(a string) (string, error) {
	x, err := parseOne(a)
	if err != nil {
		return "", err
	}
	if x < 0 {
		return "", fmt.Errorf("cannot take square root of a negative number (%s)", a)
	}
	return formatResult(math.Sqrt(x)), nil
}

// Factorial computes n! for a non-negative integer n. Returns an error for
// negatives, non-integers, or values that would overflow float64.
func Factorial(a string) (string, error) {
	x, err := parseOne(a)
	if err != nil {
		return "", err
	}
	if x < 0 {
		return "", fmt.Errorf("factorial is not defined for negative numbers")
	}
	if x != math.Trunc(x) {
		return "", fmt.Errorf("factorial is only defined for integers, got %s", a)
	}
	// 170! is the largest factorial that fits in float64; 171! overflows.
	if x > 170 {
		return "", fmt.Errorf("factorial of %s overflows float64 (max supported: 170)", a)
	}
	n := int(x)
	result := 1.0
	for i := 2; i <= n; i++ {
		result *= float64(i)
	}
	return formatResult(result), nil
}
