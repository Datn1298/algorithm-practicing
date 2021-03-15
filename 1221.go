/* 1221. Split a String in Balanced Strings
https://leetcode.com/problems/split-a-string-in-balanced-strings/
Balanced strings are those that have an equal quantity of 'L' and 'R' characters.

Given a balanced string s, split it in the maximum amount of balanced strings.

Return the maximum amount of split balanced strings.

*/

func balancedStringSplit(s string) int {
	rcount := 0
	lcount := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			rcount++
		} else {
			lcount++
		}

		if rcount == lcount {
			count++
			rcount, lcount = 0, 0
		}
	}
	return count
}