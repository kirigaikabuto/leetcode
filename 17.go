package main

import "fmt"

func letterCombinations(digits string) []string {
	alpha := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	elements := [][]string{}
	for _, v := range digits {
		for key, val := range alpha {
			if string(v) == key {
				elements = append(elements, val)
			}
		}
	}
	for _, v := range elements {
		for _, j := range v {
			fmt.Println(j)
		}
	}
	return []string{}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
