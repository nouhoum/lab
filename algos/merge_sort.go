package main

import "fmt"

/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
*/
// Looping over the second array indices
func merge(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println("nums1=", nums1)
	i := m - 1
	j := n - 1
	k := m + n - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
			k--
		} else {
			nums1[k] = nums2[j]
			j--
			k--
		}
	}

	fmt.Println("nums1=", nums1)
}

// Looping over the first array indices
func merge2(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println("nums1=", nums1)
	i := m - 1
	j := n - 1
	k := m + n - 1

	for i >= 0 {
		if j >= 0 && nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
			k--
		} else {
			nums1[k] = nums1[i]
			i--
			k--
		}
	}

	fmt.Println("nums1=", nums1)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	fmt.Println("sorting...")
	merge(nums1, m, nums2, n)
	fmt.Println("sorting...done")
}
