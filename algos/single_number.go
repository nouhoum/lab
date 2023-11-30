package main

import (
	"fmt"
)

/*
Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it.

You must implement a solution with a linear runtime complexity and use only constant extra space.
*/

func singleNumber(nums []int) int {
	freqs := make(map[int]int, 0)
	for _, n := range nums {
		freqs[n] += 1
	}

	for i, n := range freqs {
		if n == 1 {
			return i
		}
	}

	return 0
}

func main() {
	fmt.Println("singleNumber()")

	nums := []int{2, 2, 3, 2}
	checkResponse(3, nums)
}

func checkResponse(expected int, nums []int) {
	r := singleNumber(nums)
	if r != expected {
		panic(fmt.Errorf("Oops wrong response %d", r))
	}

	fmt.Printf("singleNumber() of %v is %d\n", nums, r)
}
