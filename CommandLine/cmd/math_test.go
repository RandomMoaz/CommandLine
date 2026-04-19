package cmd

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		a, b, want string
		wantErr    bool
	}{
		{"2", "3", "5", false},
		{"1.5", "2.5", "4", false},
		{"-1", "1", "0", false},
		{"abc", "1", "", true},
		{"1", "xyz", "", true},
	}
	for _, c := range cases {
		got, err := Add(c.a, c.b)
		if (err != nil) != c.wantErr {
			t.Errorf("Add(%q,%q) err=%v wantErr=%v", c.a, c.b, err, c.wantErr)
			continue
		}
		if !c.wantErr && got != c.want {
			t.Errorf("Add(%q,%q) = %q want %q", c.a, c.b, got, c.want)
		}
	}
}

func TestSubtract(t *testing.T) {
	got, err := Subtract("10", "3")
	if err != nil || got != "7" {
		t.Errorf("Subtract(10,3) = %q, %v; want 7", got, err)
	}
}

func TestMultiply(t *testing.T) {
	got, err := Multiply("6", "7")
	if err != nil || got != "42" {
		t.Errorf("Multiply(6,7) = %q, %v; want 42", got, err)
	}
}

func TestDivide(t *testing.T) {
	got, err := Divide("10", "4")
	if err != nil || got != "2.5" {
		t.Errorf("Divide(10,4) = %q, %v; want 2.5", got, err)
	}
	if _, err := Divide("1", "0"); err == nil {
		t.Errorf("Divide by zero should return error")
	}
}

func TestPower(t *testing.T) {
	got, err := Power("2", "10")
	if err != nil || got != "1024" {
		t.Errorf("Power(2,10) = %q, %v; want 1024", got, err)
	}
}

func TestMod(t *testing.T) {
	got, err := Mod("10", "3")
	if err != nil || got != "1" {
		t.Errorf("Mod(10,3) = %q, %v; want 1", got, err)
	}
	if _, err := Mod("1", "0"); err == nil {
		t.Errorf("Mod by zero should return error")
	}
}

func TestSqrt(t *testing.T) {
	got, err := Sqrt("16")
	if err != nil || got != "4" {
		t.Errorf("Sqrt(16) = %q, %v; want 4", got, err)
	}
	if _, err := Sqrt("-1"); err == nil {
		t.Errorf("Sqrt of negative should return error")
	}
}

func TestFactorial(t *testing.T) {
	cases := []struct {
		in, want string
		wantErr  bool
	}{
		{"0", "1", false},
		{"1", "1", false},
		{"5", "120", false},
		{"10", "3628800", false},
		{"-1", "", true},
		{"2.5", "", true},
		{"171", "", true},
	}
	for _, c := range cases {
		got, err := Factorial(c.in)
		if (err != nil) != c.wantErr {
			t.Errorf("Factorial(%q) err=%v wantErr=%v", c.in, err, c.wantErr)
			continue
		}
		if !c.wantErr && got != c.want {
			t.Errorf("Factorial(%q) = %q want %q", c.in, got, c.want)
		}
	}
}
