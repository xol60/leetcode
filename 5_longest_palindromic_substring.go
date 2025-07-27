package main

func longestPalindrome(s string) string {
	n := len(s)
	start, end, max := 0, 0, 1
	arr := make([][]bool, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]bool, n)
		arr[i][i] = true
	}

	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] && (arr[i-1][j+1] || j == i-1) {
				arr[i][j] = true
				if i-j+1 > max {
					start = j
					end = i
					max = i - j + 1
				}
			}
		}
	}
	return s[start : end+1]
}
