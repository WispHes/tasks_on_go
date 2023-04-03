package main

import (
	"fmt"
	"sort"
)

func main() {
	s, i := []int{23, 14, 52, 78, 33, 45, 31, 28, 77, 83, 29, 63, 60}, 60
	fmt.Println(binarySearch(s, i))
}

func binarySearch(s []int, i int) int {
	sort.Ints(s)
	fmt.Println(s)
	left, mid, right := 0, (len(s)-1)/2, len(s)-1
	for left <= right {
		if i == s[mid] {
			return s[mid]
		} else if i > s[mid] {
			left = mid + 1
		} else if i < s[mid] {
			right = mid - 1
		}
		mid = (left + right) / 2
	}
	return -1
}
