package main

import (
	"fmt"
)

func maxArea(height []int) int {
	idx := len(height) - 1
	max_w := -1

	r := 0
	for r != idx {
		tmp := (idx - r) * min(height[idx], height[r])
		if max_w < tmp {
			max_w = tmp
		}
		if height[idx] > height[r] {
			r = r + 1
		} else {
			idx = idx - 1
		}
	}
	return max_w
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {

	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(height))

	//fmt.Println(math.Min(1, 2))

}
