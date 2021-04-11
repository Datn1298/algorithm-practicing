/* 1512. Number of Good Pairs
https://leetcode.com/problems/number-of-good-pairs/

Given an array of integers nums. A pair (i,j) is called good if nums[i] == nums[j] and i < j.
Return the number of good pairs.
*/

func numIdenticalPairs(nums []int) int {
	count := 0

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}

	return count
}