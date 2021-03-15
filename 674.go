/*674. Longest Continuous Increasing Subsequence
https://leetcode.com/problems/longest-continuous-increasing-subsequence/
Given an unsorted array of integers nums, return the length of the longest continuous increasing subsequence (i.e. subarray). The subsequence must be strictly increasing.

A continuous increasing subsequence is defined by two indices l and r (l < r) such that it is [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] and for each l <= i < r, nums[i] < nums[i + 1].
*/

func findLengthOfLCIS(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sum := 1
	max := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			sum++
		} else {
			if sum > max {
				max = sum
			}
			sum = 1
		}
	}

	if sum > max {
		max = sum
	}

	return max
}