package main

import "fmt"

func containsDuplicate(nums []int) bool {
	counter := 0
	for _, v := range nums {
		counter = 0
		for _, k := range nums {
			if k == v {
				counter += 1
			}
			if counter > 1 {
				break
			}
		}
		if counter > 1 {
			break
		}
	}
	if counter > 1 {
		return true
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(arr))
}
