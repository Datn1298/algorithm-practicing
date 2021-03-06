/* 941. Valid Mountain Array
https://leetcode.com/problems/valid-mountain-array/
Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
*/

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	i := 0

	for i < len(arr)-1 && arr[i] < arr[i+1] {
		i++
	}

	if i == 0 || i == len(arr)-1 {
		return false
	}

	for i < len(arr)-1 && arr[i] > arr[i+1] {
		i++
	}
	if i == len(arr)-1 {
		return true
	}
	return false
}