package main

import (
	"fmt"
	"strconv"
)

func addBinary(a, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	output := ""
	res := 0
	for i >= 0 && j >= 0 {
		first := int(a[i] - '0')
		second := int(b[j] - '0')
		res, carry = sumBinary(first, second, carry)
		output = strconv.Itoa(res) + output
		i = i - 1
		j = j - 1
	}
	for i >= 0 {
		first := int(a[i] - '0')
		res, carry = sumBinary(first, 0, carry)
		output = strconv.Itoa(res) + output
		i = i - 1
	}
	for j >= 0 {
		second := int(b[j] - '0')
		res, carry = sumBinary(0, second, carry)
		output = strconv.Itoa(res) + output
		j = j - 1
	}
	if carry > 0 {
		output = strconv.Itoa(1) + output
	}
	return output
}

func sumBinary(a, b, carry int) (int, int) {
	res := a + b + carry
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
	fmt.Println(addBinary("111", "101"))
}
