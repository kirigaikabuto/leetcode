package main

import "fmt"

func maxProfit(prices []int) int {
	return 0
}

func getMax(nums []int) (int, int) {
	maxi := nums[0]
	maxiIndex := 0
	for i, v := range nums {
		if nums[i] >= maxi {
			maxi = v
			maxiIndex = i
		}
	}
	return maxi, maxiIndex
}

func getMin(nums []int) (int, int) {
	mini := nums[0]
	minIndex := 0
	for i, v := range nums {
		if nums[i] <= mini {
			mini = v
			minIndex = i
		}
	}
	return mini, minIndex
}

func main() {
	prices := []int{2, 4, 1}
	fmt.Println(maxProfit(prices))
}
