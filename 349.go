package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	mainNums := nums1
	secondary := nums2
	if len(nums1) > len(nums2) {
		mainNums = nums2
		secondary = nums1
	}
	for _, v := range mainNums {
		for _, k := range secondary {
			if v == k && isUnique(res, v) {
				res = append(res, v)
				break
			}
		}
	}
	return res
}

func isUnique(nums []int, k int) bool {
	isNotExist := true
	for _, v := range nums {
		if v == k {
			isNotExist = false
			break
		}
	}
	return isNotExist
}

func main() {
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
