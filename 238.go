/*238. Product of Array Except Self
https://leetcode.com/problems/product-of-array-except-self/

Given an array nums of n integers where n > 1,
return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].
*/

func productExceptSelf(nums []int) []int {

	result := make([]int, len(nums))

	preNum := 1

	for i := 0; i < len(nums); i++ {
		result[i] = preNum
		preNum *= nums[i]
	}

	preNum = 1

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= preNum
		preNum *= nums[i]
	}

	return result
}