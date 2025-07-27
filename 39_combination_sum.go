package main

func recursion_4(candidates []int, target int, now *[][]int, sum int, temp []int, start int) {
	if sum == target {
		combination := make([]int, len(temp))
		copy(combination, temp)
		*now = append(*now, combination)
		return
	}
	if sum < target {
		for i := start; i < len(candidates); i++ {
			if candidates[i]+sum <= target {
				temp = append(temp, candidates[i])
				recursion_4(candidates, target, now, sum+candidates[i], temp, i)
				temp = temp[:len(temp)-1]
			}
		}
	}
}
func combinationSum(candidates []int, target int) [][]int {
	now := [][]int{}
	recursion_4(candidates, target, &now, 0, []int{}, 0)
	return now
}
