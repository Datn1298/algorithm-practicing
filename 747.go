/* 747. Largest Number At Least Twice of Others
https://leetcode.com/problems/largest-number-at-least-twice-of-others/

In a given integer array nums, there is always exactly one largest element.

Find whether the largest element in the array is at least twice as much as every other number in the array.

If it is, return the index of the largest element, otherwise return -1.

*/

func dominantIndex(nums []int) int {
	max, res := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			res = i
		}
	}

	for i := 0; i < len(nums); i++ {
		if i != res {
			if max < nums[i]*2 {
				return -1
			}
		}
	}
	return res

}
