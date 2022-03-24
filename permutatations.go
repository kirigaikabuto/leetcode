package main

import "fmt"

func exist(el int32, arr []int32) bool {
	for _, v := range arr {
		if v == el {
			return true
		}
	}
	return false
}

func nonDivisibleSubset(k int32, s []int32) int32 {
	// Write your code here
	elements := []int32{}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if (s[i]+s[j])%k != 0 {
				if !exist(s[i], elements) {
					elements = append(elements, s[i])
				}
				if !exist(s[j], elements) {
					elements = append(elements, s[j])
				}
			}
		}
	}
	return int32(len(elements))
}

func main() {
	fmt.Println(nonDivisibleSubset(3, []int32{1, 7, 2, 4}))
}
