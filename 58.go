/* 58. Length of Last Word
https://leetcode.com/problems/length-of-last-word/

Given a string s consists of some words separated by spaces, return the length of the last word in the string. If the last word does not exist, return 0.

A word is a maximal substring consisting of non-space characters only.
*/

func lengthOfLastWord(s string) int {
	length := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && length == 0 {

		} else if s[i] != ' ' {
			length++
		} else {
			break
		}
	}

	return length
}