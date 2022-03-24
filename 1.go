package main

import "fmt"

func maxProfitTo(prices []int) int {
	mini := prices[0]
	maxi := 0
	maxiIndex := 0
	miniIndex := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] > maxi && miniIndex < i {
			maxi = prices[i]
		}
		if prices[i] < mini && maxiIndex < i {
			mini = prices[i]
		}
	}
	return maxi - mini
}

func main() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfitTo(prices))
}
