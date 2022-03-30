package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	res := ""
	i := len(num1) - 1
	j := len(num2) - 1
	sum := 0
	c := 0
	for i >= 0 && j >= 0 {
		first := int(num1[i] - '0')
		second := int(num2[j] - '0')
		sum, c = stringsSum(first, second, c)
		res = strconv.Itoa(sum) + res
		i -= 1
		j -= 1
	}
	for i >= 0 {
		first := int(num1[i] - '0')
		sum, c = stringsSum(first, 0, c)
		res = strconv.Itoa(sum) + res
		i -= 1
	}
	for j >= 0 {
		second := int(num2[j] - '0')
		sum, c = stringsSum(0, second, c)
		res = strconv.Itoa(sum) + res
		j -= 1
	}
	if c > 0 {
		res = strconv.Itoa(c) + res
	}
	return res
}

func stringsSum(a, b, c int) (int, int) {
	sum := a + +b + c
	sums := strconv.Itoa(sum)
	if len(sums) != 2 {
		return sum, 0
	}
	n0, _ := strconv.Atoi(string(sums[0]))
	n1, _ := strconv.Atoi(string(sums[1]))
	return n1, n0
}

func main() {
	fmt.Println(addStrings("456", "77"))
}
