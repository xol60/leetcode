package main

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func maxArea(height []int) int {
	start, end := 0, len(height)-1
	max := 0
	for start < end {
		max = maxInt(max, (end-start)*minInt(height[start], height[end]))
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return max
}
