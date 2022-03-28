package main

import (
	"fmt"
	"strconv"
)

func addTwoBinary(a, b string) string {
	var res string
	var temp int
	var c int
	i := len(a) - 1
	j := len(b) - 1
	for i >= 0 && j >= 0 {
		first := int(a[i] - '0')
		second := int(b[j] - '0')
		temp, c = check(first, second, c)
		res = strconv.Itoa(temp) + res
		i -= 1
		j -= 1
	}
	for i >= 0 {
		first := int(a[i] - '0')
		temp, c = check(first, 0, c)
		res = strconv.Itoa(temp) + res
		i -= 1
	}
	for j >= 0 {
		second := int(b[j] - '0')
		temp, c = check(0, second, c)
		res = strconv.Itoa(temp) + res
		j -= 1
	}
	if c != 0 {
		res = "1" + res
	}
	return res
}

func check(a, b, c int) (int, int) {
	res := a + b + c
	if res == 0 {
		return 0, 0
	}
	if res == 1 {
		return 1, 0
	}
	if res == 2 {
		return 0, 1
	}
	if res == 3 {
		return 1, 1
	}
	return 0, 0
}

func main() {
	s1 := "1010"
	s2 := "1011"
	fmt.Println(addTwoBinary(s1, s2))
}
