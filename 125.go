package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s1 := ""
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			s1 += strings.ToLower(string(v))
		}
	}
	for i, _ := range s1 {
		if s1[i] != s1[len(s1)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	s := "0P"
	fmt.Println(isPalindrome(s))
}
