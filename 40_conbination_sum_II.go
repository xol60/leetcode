package main

import (
	"sort"
)

func recursion_5(candidates []int, target int, now *[][]int, sum int, temp []int, start int) {
	if sum == target {
		combination := make([]int, len(temp))
		copy(combination, temp)
		*now = append(*now, combination)
		return
	}
	if sum < target {
		for i := start + 1; i < len(candidates); i++ {
			if i == start+1 || (i > start+1 && candidates[i] != candidates[i-1]) {
				if candidates[i]+sum <= target {
					temp = append(temp, candidates[i])
					recursion_5(candidates, target, now, sum+candidates[i], temp, i)
					temp = temp[:len(temp)-1]
				}
			}
		}
	}
}
func combinationSum2(candidates []int, target int) [][]int {
	now := [][]int{}
	sort.Ints(candidates)
	//fmt.Println(candidates)
	recursion_5(candidates, target, &now, 0, []int{}, -1)
	return now
}
