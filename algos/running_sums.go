package main

import "fmt"

func runningSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	return sums
}

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(runningSum(arr))
}
