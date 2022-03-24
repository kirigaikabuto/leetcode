package main

import "fmt"

func validPalindrome(s string) bool {
	if len(s) == 1 {
		return false
	}
	n := len(s) - 1
	for i, _ := range s {
		if s[i] != s[n-i] {
			newS := ""
			if i == 0 {
				newS = s[:n]
			} else {
				newS = s[:n-i] + s[n-i+1:]
			}
			fmt.Println(newS)
			return validPalindrome(newS)
		}
	}
	return true
}

func main() {
	s2 := "deeee"
	fmt.Println(validPalindrome(s2))
}
