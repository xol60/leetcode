package main

import "math"

func reverse(x int) int {
	if x < 0 {
		return (-1) * reverse((-1)*x)
	}
	var r int = 0
	for x > 0 {
		// fmt.Println(x/10)
		if r > math.MaxInt32/10 {
			return 0
		}
		r *= 10
		r += x % 10
		x /= 10
	}
	return r
}
