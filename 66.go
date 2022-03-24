package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func plusOne(digits []int) []int {
	number := new(big.Int)
	el := new(big.Int).Set(big.NewInt(1))
	digitsStr := ""
	for _, v := range digits {
		digitsStr += strconv.Itoa(v)
	}
	number, _ = number.SetString(digitsStr, 10)
	number.Add(number, el)
	newString := number.String()
	newDigits := []int{}
	for _, v := range newString {
		newEl, _ := strconv.Atoi(string(v))
		newDigits = append(newDigits, newEl)
	}
	return newDigits
}

func main() {
	arr := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	response := plusOne(arr)
	fmt.Println(response)
}
