package main
func lengthOfLongestSubstring(s string) int {
	start, max, n := 0, 0, len(s)
	maps := make(map[byte]int)
	for end := 0; end < n; end++ {
		for start <= end {
			v, ok := maps[s[end]]
			if !ok || v == -1 {
				maps[s[end]] = end
				break
			} else {
				maps[s[start]] = -1
				start++
			}
		}
		max = maxInt(max, end-start+1)
	}
	return max
}
