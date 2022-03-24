package main

import (
	"fmt"
	"strconv"
)

func join(nums []int) int64 {
	var str string
	for i := range nums {
		str += strconv.Itoa(nums[i])
	}
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	} else {
		return num
	}
}

func multiply(num1 string, num2 string) string {
	constant := 48
	firstVal := []int{}
	secondVal := []int{}
	for _, v := range num1 {
		firstVal = append(firstVal, int(v)-constant)
	}
	for _, v := range num2 {
		secondVal = append(secondVal, int(v)-constant)
	}
	return strconv.Itoa(join(firstVal) * join(secondVal))
}

func main() {
	fmt.Println(multiply("160", "4"))
}
