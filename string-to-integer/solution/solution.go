package solution

import "fmt"

import "strconv"

const (
	lowerboundsErr = "passed lower bounds"
	upperboundsErr = "passed upper bounds"

	minValue = -2147483648
)

func myAtoi(str string) int {
	n, err := parse(str)
	if err != nil {
		switch err.Error() {
		case lowerboundsErr:
			return -2147483648
		case upperboundsErr:
			return 2147483647
		}
	}
	return n
}

func parse(s string) (int, error) {
	bytes := []byte(s)
	record := false
	negative := false
	var number int32
	fmt.Println(s, bytes)
	for i, r := range bytes {
		fmt.Printf("pos %d holds '%s' or '%b'\n", i, string(r), r)
		if r|0x20 == 0x20 {
			fmt.Println("hit space")
			if record {
				break
			}
			continue
		}
		if (int(r) >= 0 && int(r) < 10) || r|0x2d == 0x2d {
			record = true
		}
		if !record {
			continue
		}
		if record {
			fmt.Println("in record")
			if r|0x2d == 0x2d {
				negative = true
			}
			add, err := strconv.Atoi(string(r))
			if err != nil {
				break
			}
			// multiply per iteration
			number = (number << 3) + (number << 1)
			number += int32(add)
		}
	}
	if negative {
		return -1, nil
	}
	return int(number), nil
}
