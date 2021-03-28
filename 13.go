/* 13. Roman to Integer
https://leetcode.com/problems/roman-to-integer/

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.
*/

func romanToInt(s string) int {
	res := charToInt(s[len(s)-1])
	for i := len(s) - 2; i >= 0; i-- {
		if charToInt(s[i]) < charToInt(s[i+1]) {
			res -= charToInt(s[i])
		} else {
			res += charToInt(s[i])
		}
	}
	return res
}

func charToInt(s byte) int {
	if s == 'I' {
		return 1
	}
	if s == 'V' {
		return 5
	}
	if s == 'X' {
		return 10
	}
	if s == 'L' {
		return 50
	}
	if s == 'C' {
		return 100
	}
	if s == 'D' {
		return 500
	}
	if s == 'M' {
		return 1000
	}
	return 0
}
