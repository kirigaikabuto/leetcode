package main

import (
	"fmt"
	"math"
	"strconv"
)

func isHappy(n int) bool {
	for true {
		s := strconv.Itoa(n)
		if n == 1 {
			return true
		}
		if len(s) == 1 {
			return false
		}
		n = calc(n)
	}
	return false
}

func calc(n int) int {
	s := strconv.Itoa(n)
	sumi := 0.0
	for _, v := range s {
		num, _ := strconv.Atoi(string(v))
		sumi = sumi + math.Pow(float64(num), 2.0)
	}
	return int(sumi)
}

func main() {
	n := 2
	fmt.Println(isHappy(n))

}
