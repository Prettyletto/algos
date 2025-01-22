package main

// https://leetcode.com/problems/two-sum/submissions/1517370468/
func twoSums(nums []int, target int) []int {
	candidates := map[int]int{}
	results := make([]int, 2, 2)

	for i := range nums {

		c := target - nums[i]
		v, ok := candidates[nums[i]]
		if ok {
			results[0] = v
			results[1] = i

		}
		candidates[c] = i
	}

	return results
}
