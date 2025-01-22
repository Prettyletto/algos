package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
func removeDuplicates(nums []int) int {
	k := 1
	nums[k] = nums[0]
	k++

	for i := range nums {
		if nums[k-1] != nums[i] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
