package main

import (
	"fmt"
	"math/big"
)

func extraLongFactorials(n int32) {
	// Write your code here
	fact := big.NewInt(1)
	var i int32
	for i = 1; i <= n; i++ {
		bInt := big.NewInt(int64(i))
		fact = fact.Mul(fact, bInt)
	}
	fmt.Println(fact)
}

func main() {
	extraLongFactorials(3)
}
