/* 1464. Maximum Product of Two Elements in an Array
https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/

Given the array of integers nums, you will choose two different indices i and j of that array.
Return the maximum value of (nums[i]-1)*(nums[j]-1).
*/

func maxProduct(nums []int) int {
	sortArray(nums)

	if (nums[0]-1)*(nums[1]-1) > (nums[len(nums)-1]-1)*(nums[len(nums)-2]-1) {
		return (nums[0] - 1) * (nums[1] - 1)
	} else {
		return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
	}

	return 0
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