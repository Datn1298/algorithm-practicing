/* 792. Number of Matching Subsequences
https://leetcode.com/problems/number-of-matching-subsequences/

Given a string s and an array of strings words, return the number of words[i] that is a subsequence of s.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".
*/

func numMatchingSubseq(s string, words []string) int {
	count := 0
	i := 0
	for i < len(words) {
		if checkSubstring(s, words[i]) {
			count++
		}
		i++
	}

	return count
}

func checkSubstring(a, b string) bool {
	// b is substring a??

	iterator := 0

	count := 0
	i := 0
	for i < len(b) {

		for j := iterator; j < len(a); j++ {
			if a[j] == b[i] {
				iterator = j + 1
				count++
				break
			}
		}
		i++
	}

	if count < len(b) {
		return false
	}
	return true

}
