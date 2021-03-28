/* 1646. Get Maximum in Generated Array
https://leetcode.com/problems/get-maximum-in-generated-array/

You are given an integer n. An array nums of length n + 1 is generated in the following way:

nums[0] = 0
nums[1] = 1
nums[2 * i] = nums[i] when 2 <= 2 * i <= n
nums[2 * i + 1] = nums[i] + nums[i + 1] when 2 <= 2 * i + 1 <= n
Return the maximum integer in the array nums​​​.
*/

func getMaximumGenerated(n int) int {
	max := 0

	nums := make([]int, n+1)

	for i := 0; i <= n; i++ {
		if i == 0 {
			nums[i] = 0
			if nums[i] > max {
				max = nums[i]
			}
		} else if i == 1 {
			nums[i] = 1
			if nums[i] > max {
				max = nums[i]
			}
		} else if i%2 == 0 {
			nums[i] = nums[i/2]
			if nums[i] > max {
				max = nums[i]
			}
		} else {
			nums[i] = nums[i/2] + nums[i/2+1]
			if nums[i] > max {
				max = nums[i]
			}
		}
	}

	return max
}