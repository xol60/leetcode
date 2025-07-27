package main

func isValid(s string) bool {
	stack := []rune{}
	par := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, v := range s {
		value, exist := par[v]
		if len(stack) == 0 {
			if exist {
				return false
			}
			stack = append(stack, v)
		} else {
			if exist {
				if value == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else {
				stack = append(stack, v)
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
