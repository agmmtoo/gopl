package main

import (
	"fmt"
)

func main() {
	m := map[string]int{}
	fmt.Printf("len: %d, nil: %t\n", len(m), m == nil)
	m["carol"] = 21
	fmt.Println(m)
}
