package main

import "fmt"

func isAlienSorted(words []string, order string) bool {
	alpha := make(map[string]int)
	for i, v := range order {
		alpha[string(v)] = i
	}
	for i := 0; i < len(words)-1; i++ {
		for j := 0; i < len(words[i]); j++ {
			if j >= len(words[i+1]) {
				return false
			}
			if words[i][j] != words[i+1][j] {
				if alpha[string(words[i][j])] > alpha[string(words[i+1][j])] {
					return false
				}
				break
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
}
