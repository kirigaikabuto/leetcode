package main

import "fmt"

//func maxSubArray(nums []int) int {
//	maxSum := 0
//	for _, v := range nums {
//		sumi := 0
//		for _, k := range nums {
//			sumi +=
//		}
//	}
//}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	n := len(nums)
	middle := n / 2
	fmt.Println(nums[:middle])
	fmt.Println(nums[middle:])
	middle = middle / 3
	fmt.Println(nums[:middle])
	fmt.Println(nums[middle : middle])
}
