package main

import "fmt"

func maxProfit(prices []int) int {
	result := 0
	mini := prices[0]
	maxi := 0
	miniIndex := 0
	for i, v := range prices {
		if v < mini {
			mini = v
			miniIndex = i
		}
	}
	if miniIndex == len(prices)-1 {
		podMini := prices[0]
		podMiniIndex := 0
		prices = prices[:miniIndex]
		fmt.Println(prices)
		for i, v := range prices {
			if v < podMini {
				podMini = v
				podMiniIndex = i
			}
		}
		fmt.Println(podMini)
		fmt.Println(podMiniIndex)
		if podMiniIndex == len(prices)-1 {
			return 0
		}
		for j := podMiniIndex + 1; j < len(prices); j++ {
			if prices[j] > maxi {
				maxi = prices[j]
			}
		}
		if podMini == 0 && maxi != 0 {
			return maxi + 1
		}
		if maxi == 0 {
			return 0
		}
		return maxi - podMini
	}
	for j := miniIndex + 1; j < len(prices); j++ {
		if prices[j] > maxi {
			maxi = prices[j]
		}
	}
	if mini == 0 && maxi != 0 {
		return maxi + 1
	}
	result = maxi - mini
	return result
}

func main() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}
