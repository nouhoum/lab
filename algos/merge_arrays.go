package main

import "fmt"

// n = lenght of nums2
// m = length of nums1
// Time O(n+m) | Space O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := n - 1
	j := m - 1
	k := n + m - 1 // the last element of th combined array
	for i >= 0 {
		if j >= 0 && nums1[j] > nums2[i] {
			nums1[k] = nums1[j]
			k--
			j--
		} else {
			nums1[k] = nums2[i]
			k--
			i--
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{2, 5, 6}, 3)

	fmt.Println(nums1)

	nums1 = []int{0}
	nums2 := []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Println(nums1)
}
