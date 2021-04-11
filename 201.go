/* 201. Bitwise AND of Numbers Range
https://leetcode.com/problems/bitwise-and-of-numbers-range/

Given two integers left and right that represent the range [left, right], return the bitwise AND of all numbers in this range, inclusive.
*/

func rangeBitwiseAnd(left int, right int) int {
	count := 0
	for left != 0 && left != right {
		count++
		left /= 2
		right /= 2
	}

	return left << count
}