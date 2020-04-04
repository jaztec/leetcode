package solution

import (
	"errors"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	type test struct {
		in  string
		out int
	}
	tests := []test{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", minValue},
		{"191283472332", maxValue},
		{"+1", 1},
		{"+4193 with words", 4193},
		{"words and +987", 0},
		{"+-2", 0},
		{"-+2", 0},
		{"010", 10},
		{"   +0 123", 0},
		{"2147483648", maxValue},
		{"2147483646", 2147483646},
		{"-2147483649", minValue},
		{"0-1", 0},
		{"0  123", 0},
		{"  - 12", 0},
		{" ++1", 0},
		{" --1", 0},
	}
	for _, tN := range tests {
		if a := myAtoi(tN.in); a != tN.out {
			t.Errorf("got %d from '%s' but wanted %d", a, tN.in, tN.out)
		}
	}
}

func TestParse(t *testing.T) {
	type test struct {
		in  string
		out int
		err error
	}
	tests := []test{
		{"42", 42, nil},
		{"   -42", -42, nil},
		{"4193 with words", 4193, nil},
		{"words and 987", 0, nil},
		{"2147483646", 2147483646, nil},
		{"-91283472332", 0, errors.New(lowerboundsErr)},
		{"191283472332", 0, errors.New(upperboundsErr)},
		{"2147483648", 0, errors.New(upperboundsErr)},
		{"-2147483649", 0, errors.New(lowerboundsErr)},
	}
	for _, tN := range tests {
		a, err := parse(tN.in)
		if a != tN.out {
			t.Errorf("got %d from '%s' but wanted %d", a, tN.in, tN.out)
		}
		if (err == nil && tN.err != nil) || (err != nil && &tN.err == nil) {
			t.Errorf("got '%v' from '%s' but wanted '%v'", err, tN.in, tN.err)
		}
		if err != nil && tN.err != nil && err.Error() != tN.err.Error() {
			t.Errorf("got %d and '%v' from '%s' but wanted %d and '%v'", a, err, tN.in, tN.out, tN.err)
		}
	}
}

func BenchmarkMyAtoi(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		myAtoi("10")
	}
}
