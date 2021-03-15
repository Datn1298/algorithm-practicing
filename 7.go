/*7. Reverse Integer
https://leetcode.com/problems/reverse-integer/

Given a signed 32-bit integer x, return x with its digits reversed.
If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

*/

func reverse(x int) int {
	rev_int := 0

	for x != 0 {
		rev_int = rev_int*10 + x%10
		x /= 10
	}

	if rev_int > (1<<31 - 1) {
		return 0
	} else if rev_int < -1<<31 {
		return 0
	} else {
		return rev_int
	}

}