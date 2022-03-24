package main

import "fmt"

func majorityElement(nums []int) int {
	finalCount := 0
	finaValue := 0
	for _, v := range nums {
		count := 0
		for _, k := range nums {
			if v == k {
				count++
			}
		}
		if count > finalCount {
			finalCount = count
			finaValue = v
		}
	}
	return finaValue
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
