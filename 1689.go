/* 1689. Partitioning Into Minimum Number Of Deci-Binary Numbers
https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/

A decimal number is called deci-binary if each of its digits is either 0 or 1 without any leading zeros. For example, 101 and 1100 are deci-binary, while 112 and 3001 are not.

Given a string n that represents a positive decimal integer, return the minimum number of positive deci-binary numbers needed so that they sum up to n.
*/

func minPartitions(n string) int {
	max := 0

	for i := 0; i < len(n); i++ {
		// if convert(n[i]) > max {
		//     max = convert(n[i])
		// }

		if int(n[i]-'0') > max {
			max = int(n[i] - '0')
		}
	}
	return max
}