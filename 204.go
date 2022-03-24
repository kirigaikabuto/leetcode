package main

import "fmt"

func countPrimes(n int) int {
	globalCounter := 0
	for i := 1; i < n; i++ {
		if getPrime(i) {
			globalCounter += 1
		}
	}
	return globalCounter
}

func getPrime(number int) bool {
	if number == 2 {
		return true
	}
	if number < 2 {
		return false
	}
	if isOdd(number) {
		for i := 3; i <= number/3; i += 2 {
			if number%i == 0 {
				return false
			}
		}
		return true
	}
	return false
}

func isOdd(number int) bool {
	if number%2 != 0 {
		return true
	}
	return false
}

func main() {
	n := 10
	fmt.Println(countPrimes(n))
}
