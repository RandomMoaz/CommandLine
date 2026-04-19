package cmd

import "testing"

func TestEvaluate(t *testing.T) {
	cases := []struct {
		expr, want string
	}{
		{"1+1", "2"},
		{"2 + 3 * 4", "14"},          // precedence
		{"(2 + 3) * 4", "20"},        // parens
		{"10 - 2 - 3", "5"},          // left-associative
		{"2 ^ 10", "1024"},           // power
		{"2 ^ 3 ^ 2", "512"},         // right-associative power: 2^(3^2)=2^9=512
		{"-5 + 10", "5"},             // unary minus
		{"--5", "5"},                 // double unary minus
		{"10 % 3", "1"},              // modulo
		{"1.5 + 2.5", "4"},           // decimals
		{"1e3 + 1", "1001"},          // scientific notation
		{"2 * (3 + (4 - 1))", "12"},  // nested parens
		{"-(2 + 3)", "-5"},           // unary on parenthesized expr
	}
	for _, c := range cases {
		got, err := Evaluate(c.expr)
		if err != nil {
			t.Errorf("Evaluate(%q) error: %v", c.expr, err)
			continue
		}
		if got != c.want {
			t.Errorf("Evaluate(%q) = %q want %q", c.expr, got, c.want)
		}
	}
}

func TestEvaluateErrors(t *testing.T) {
	bad := []string{
		"",
		"1 +",
		"(1 + 2",
		"1 + 2)",
		"1 / 0",
		"5 % 0",
		"abc",
		"1 + + ",
		"1 2",   // two numbers with no operator -> trailing input error
	}
	for _, expr := range bad {
		if _, err := Evaluate(expr); err == nil {
			t.Errorf("Evaluate(%q) expected error, got none", expr)
		}
	}
}
