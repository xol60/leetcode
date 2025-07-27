package main
func twoSum(nums []int, target int) []int {
    temp := make (map[int]int)
	for i:=0;i<len(nums);i++{
		if j,ok:=temp[target-nums[i]];ok{
			return []int{i,j}
		} else {
			temp[nums[i]]=i
		}
	}
	return []int{}
}