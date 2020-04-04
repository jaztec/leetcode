package solution

import (
	"errors"
	"math"
	"strconv"
)

const (
	lowerboundsErr = "passed lower bounds"
	upperboundsErr = "passed upper bounds"

	minValue = -2147483648
	maxValue = 2147483647
)

func myAtoi(str string) int {
	n, err := parse(str)
	if err != nil {
		switch err.Error() {
		case lowerboundsErr:
			return minValue
		case upperboundsErr:
			return maxValue
		default:
			panic(err)
		}
	}
	return n
}

func parse(s string) (int, error) {
	bytes := []byte(s)
	record := 0
	ticks := 0
	negative := false
	minusSet := false
	plusSet := false
	var number int32
	for _, b := range bytes {
		// space
		if b|0x20 == 0x20 {
			if ticks > 0 || plusSet || minusSet {
				break
			}
			continue
		}
		// minus
		if b|0x2d == 0x2d && record == 0 && ticks == 0 {
			if !plusSet && !minusSet {
				negative = true
				minusSet = true
				continue
			}
			break
		}
		// plus
		if b|0x2b == 0x2b && record == 0 && ticks == 0 {
			if !minusSet && !plusSet {
				plusSet = true
				continue
			}
			break
		}

		add, err := strconv.Atoi(string(b))
		if err != nil {
			break
		}
		if add == 0 && record == 0 {
			ticks++
			continue
		}

		temp := int64(number)
		if record > 0 {
			temp *= 10
		}

		if temp+int64(add) > maxValue || temp+int64(add) == minValue {
			if negative {
				return 0, errors.New(lowerboundsErr)
			}
			return 0, errors.New(upperboundsErr)
		}
		number = int32(temp)

		number += int32(add)
		record++
		ticks++
	}
	if negative {
		number = -int32(math.Abs(float64(number)))
	} else {
		if number == minValue {
			return 0, errors.New(upperboundsErr)
		}
		number = int32(math.Abs(float64(number)))
	}
	return int(number), nil
}
