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
		{"-91283472332", -2147483648},
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
		{"-91283472332", 0, errors.New("passed lower bounds")},
		{"191283472332", 0, errors.New("passed upper bounds")},
	}
	for _, tN := range tests {
		if a, err := parse(tN.in); a != tN.out || err != tN.err {
			t.Errorf("got %d and %v from %s but wanted %d and %v", a, err, tN.in, tN.out, tN.err)
		}
	}
}
