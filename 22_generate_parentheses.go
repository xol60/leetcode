package main

func recursion_2(arr []string, open, close int, s string) []string {
	if close == 0 {
		arr = append(arr, s)
	}
	if open > 0 {
		arr = recursion_2(arr, open-1, close, s+"(")
	}
	if close > open {
		arr = recursion_2(arr, open, close-1, s+")")
	}
	//fmt.Println(arr)
	return arr
}
func generateParenthesis(n int) []string {
	arr := []string{}
	return recursion_2(arr, n, n, "")
}
