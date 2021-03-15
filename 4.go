/*4. Median of Two Sorted Arrays
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	result := []int{}
	l, r := 0, 0

	for l < len(nums1) && r < len(nums2) {
		if nums1[l] < nums2[r] {
			result = append(result, nums1[l])
			l++
		} else {
			result = append(result, nums2[r])
			r++
		}
	}

	if l == len(nums1) {
		for r < len(nums2) {
			result = append(result, nums2[r])
			r++
		}
	}

	if r == len(nums2) {
		for l < len(nums1) {
			result = append(result, nums1[l])
			l++
		}
	}

	if len(result)%2 != 0 {
		return float64(result[len(result)/2])
	} else {
		return float64(result[len(result)/2-1]+result[len(result)/2]) / 2
	}
}
