/* 12. Integer to Roman
https://leetcode.com/problems/integer-to-roman/

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral.
*/

func intToRoman(num int) string {
	m := [4]string{"", "M", "MM", "MMM"}
	c := [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	x := [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	i := [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	thousands := m[num/1000]
	hundreds := c[(num%1000)/100]
	tens := x[(num%100)/10]
	unit := i[num%10]

	return thousands + hundreds + tens + unit
}