package main

func change(amount int, coins []int) int {
	//n:=len(coins)
	arr := make([]int, amount+1)
	arr[0] = 1
	for _, v := range coins {
		for i := v; i <= amount; i++ {
			arr[i] += arr[i-v]
		}
	}
	return arr[amount]
}
