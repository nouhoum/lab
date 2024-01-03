package main

import (
	"fmt"
)

func minimumOperations(nums []int) int {
	dict := map[int]bool{}
	for _, n := range nums {
		if n != 0 {
			dict[n] = true
		}
	}

	count := 0
	for _, _ = range dict {
		count++
	}

	return count
}

func minimumOperations2(nums []int) int {
	fmt.Println("init nums=", nums)
	if len(nums) == 1 {
		if nums[0] == 0 {
			return 0
		} else {
			return 1
		}
	}

	x := nums[0]
	for _, n := range nums {
		if x == 0 {
			x = n
		} else if x > n && n > 0 {
			x = n
		}
	}

	fmt.Println("init x=", x)
	nonZero := true
	incCount := false
	count := 0
	for nonZero {
		nonZero = false
		min := 100
		incCount = false
		for i := 0; i < len(nums); i++ {
			curr := nums[i]
			if curr > 0 {
				incCount = true
				newElement := curr - x
				nums[i] = newElement

				if newElement > 0 && min >= newElement {
					min = newElement
				}
				if newElement > 0 {
					nonZero = true
				}
			}
		}

		if incCount {
			count++
		}
		x = min
	}

	return count
}

func main() {
	// nums := []int{1, 5, 0, 3, 5}
	// fmt.Println(minimumOperations(nums))
	// nums2 := []int{0}
	// // fmt.Println(minimumOperations(nums2))
	// nums3 := []int{2}
	// fmt.Println(minimumOperations(nums3))
	// nums3 = []int{0, 0, 0, 0}
	// fmt.Println(minimumOperations(nums3) == 0)

	nums3 := []int{1, 5, 0, 3, 5}
	ef := minimumOperations(nums3)
	fmt.Println("ef=", ef, ef == 3)

	nums3 = []int{7, 28, 34, 76}
	e := minimumOperations(nums3)
	fmt.Println("e=", e, e == 4)

	nums3 = []int{1, 2, 3, 5}
	e = minimumOperations(nums3)
	fmt.Println("e=", e, e == 4)

	nums3 = []int{0, 47, 91, 61, 60}
	e = minimumOperations(nums3)
	fmt.Println("e=", e, e == 4)
}
