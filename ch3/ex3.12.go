package main

import (
	"fmt"
)

func isAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, sortedS1 := range s1 {

		_, a := fmt.Println(sortedS1)
		fmt.Println(a)
	}

	return true
}

func main() {
	isAnagrams("lmao", "omal")
}
