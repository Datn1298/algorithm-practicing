/* 1550. Three Consecutive Odds
https://leetcode.com/problems/three-consecutive-odds/

Given an integer array arr, return true if there are three consecutive odd numbers in the array. Otherwise, return false.
*/

func threeConsecutiveOdds(arr []int) bool {
	count := 0
	i := 0
	for i < len(arr)-2 {
		if arr[i]%2 == 1 {
			if (arr[i+1]%2 == 1) && arr[i+2]%2 == 1 {
				count++
			}
		}
		i++
	}

	if count > 0 {
		return true
	}
	return false

}
