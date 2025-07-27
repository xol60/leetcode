package main

func search(nums []int, target int) int {
	mid := 0
	left, right := 0, len(nums)-1
	for left <= right {
		mid = (left + right) / 2
		//fmt.Println(left,right)
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			if (nums[left] > nums[mid]) || (nums[left] <= target) {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if (nums[right] < nums[mid]) || (target <= nums[right]) {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
