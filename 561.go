/* 561. Array Partition I
https://leetcode.com/problems/array-partition-i/
Given an integer array nums of 2n integers, group these integers into n pairs (a1, b1), (a2, b2), ..., (an, bn) such that the sum of min(ai, bi) for all i is maximized. Return the maximized sum.
*/

func arrayPairSum(nums []int) int {
	sortArray(nums)

	sum := 0

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			sum += nums[i]
		}
	}

	return sum
}

func sortArray(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}