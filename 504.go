/* 504. Base 7
https://leetcode.com/problems/base-7/

Given an integer, return its base 7 string representation.
*/

func convertToBase7(num int) string {
	if num == 0 {
		return string('0')
	}

	if num < 0 {
		return string('-') + convertToBase7(-num)
	}

	str := ""

	for num != 0 {
		str = strconv.Itoa(num%7) + str
		num /= 7
	}

	return str
}
