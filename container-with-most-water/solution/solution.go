package solution

type weight struct {
	pos    int
	target int
	score  int
}

func maxArea(height []int) int {
	w := weight{}
	for pos := 0; pos < len(height); pos++ {
		for target := pos; target < len(height); target++ {
			mul := height[pos]
			if height[pos] > height[target] {
				mul = height[target]
			}
			score := mul * (target - pos)

			if score > w.score {
				w = weight{
					pos:    pos,
					target: target,
					score:  mul * (target - pos),
				}
			}
		}
	}
	return w.score
}
