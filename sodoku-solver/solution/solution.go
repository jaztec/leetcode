package solution

import (
	"errors"
	"strings"
)

var charMap = map[byte]uint8{
	46: 0,
	49: 1,
	50: 2,
	51: 3,
	52: 4,
	53: 5,
	54: 6,
	55: 7,
	56: 8,
	57: 9,
}

var intMap = map[uint8]byte{
	0: 46,
	1: 49,
	2: 50,
	3: 51,
	4: 52,
	5: 53,
	6: 54,
	7: 55,
	8: 56,
	9: 57,
}

type boardType [][]byte

type pos struct {
	x, y uint8
}

func (bt boardType) String() string {
	var builder strings.Builder
	builder.Grow(100)
	for i, b := range bt {
		// format line
		b = append(b, []byte("00")...)
		for j := range b {
			if b[j] < 40 {
				b[j] = intMap[b[j]]
			}
		}
		for _, v := range [2]int{3, 7} {
			copy(b[v+1:], b[v:])
			b[v] = byte(32)
		}
		builder.Write(b)
		builder.Write([]byte("\n"))
		if (i+1)%3 == 0 {
			builder.Write([]byte("\n"))
		}
	}
	return builder.String()
}

func solveSudoku(board boardType) {
	err := solve(board)
	if err != nil {
		panic(err)
	}
}

func solve(board boardType) error {
	for row, i := range board {
		for col, val := range i {
			if val == 46 {
				po := pos{uint8(col), uint8(row)}
				solved := false
				for x := 0; x < 9; x++ {
					if validNumber(board, po, x+1) {
						board[row][col] = intMap[uint8(x+1)]
						if err := solve(board); err != nil {
							board[row][col] = byte(46)
						} else {
							solved = true
							break
						}
					}
				}
				if !solved {
					return errors.New("No valid number found")
				}
			}
		}
	}
	return nil
}

func validNumber(board boardType, p pos, n int) bool {
	set := make([]uint8, 9)
	// get row
	for i, v := range board[p.y] {
		if uint8(i) == p.x {
			set[i] = uint8(n)
		} else {
			set[i] = charMap[v]
		}
	}
	if valid := validSet(set); !valid {
		return false
	}

	// get col
	for i, row := range board {
		if uint8(i) == p.y {
			set[i] = uint8(n)
		} else {
			set[i] = charMap[row[p.x]]
		}
	}
	if valid := validSet(set); !valid {
		return false
	}

	// get block
	for i, po := range block(p) {
		if po == p {
			set[i] = uint8(n)
		} else {
			set[i] = charMap[board[po.y][po.x]]
		}
	}
	if valid := validSet(set); !valid {
		return false
	}

	return true
}

// validSet checks 8 values in a set if the number already occurs here
func validSet(set []uint8) bool {
	matched := make(map[uint8]struct{})
	var found struct{}
	for _, c := range set {
		if c == 0 {
			continue
		}
		if _, ok := matched[c]; ok {
			return false
		}
		matched[c] = found
	}
	return true
}

func block(p pos) []pos {
	po := make([]pos, 0, 9)
	minY := int(p.y/3) * 3
	minX := int(p.x/3) * 3
	for y := minY; y < minY+3; y++ {
		for x := minX; x < minX+3; x++ {
			po = append(po, pos{uint8(x), uint8(y)})
		}
	}
	return po
}
