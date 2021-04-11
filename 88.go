/* 88. Merge Sorted Array
https://leetcode.com/problems/merge-sorted-array/

Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
The number of elements initialized in nums1 and nums2 are m and n respectively. You may assume that nums1 has a size equal to m + n such that it has enough space to hold additional elements from nums2.
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	for k := m + n - 1; k >= 0; k-- {
		if (i >= 0 && j >= 0 && nums1[i] > nums2[j]) || (j < 0) {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
