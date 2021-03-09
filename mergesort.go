package main

import (
	"fmt"
)

func merge(left, right []int) []int {
	res := make([]int, len(left)+len(right))

	i := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			res[i] = right[0]
			right = right[1:]
		} else {
			res[i] = left[0]
			left = left[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		res[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		res[i] = right[j]
		i++
	}

	return res
}

func mergeSort(list []int) []int {
	length := len(list)

	if length == 1 {
		return list
	}

	middle := int(length / 2)

	left := make([]int, middle)
	right := make([]int, length-middle)

	for i := 0; i < middle; i++ {
		left[i] = list[i]
	}

	for i := middle; i < length; i++ {
		right[i-middle] = list[i]
	}

	return merge(left, right)
}

func main() {
	nums := []int{1, 4, 5, 2, 5, 7, 8}

	fmt.Println(mergeSort(nums))
}
