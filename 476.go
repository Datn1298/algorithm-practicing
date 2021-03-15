/*476. Number Complement
https://leetcode.com/problems/number-complement/
Given a positive integer num, output its complement number. The complement strategy is to flip the bits of its binary representation.
*/

func findComplement(num int) int {
	rel := 0

	for i := 1; ; i++ {
		rel += (1 - num%2) * Pow(2, i-1)
		num /= 2

		if num == 0 {
			return rel
		}
	}
}

func Pow(a, b int) int {
	if b == 0 {
		return 1
	} else if b%2 == 0 {
		return (Pow(a, b/2)) * (Pow(a, b/2))
	} else {
		return a * (Pow(a, b/2)) * (Pow(a, b/2))
	}

}
