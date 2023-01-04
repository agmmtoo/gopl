package main

import (
	"fmt"
)

func squares() (s func() int) {
	var x int
	s = func() int {
		x++
		return x * x
	}
	return
	
}

func main() {
	f := squares()
	fmt.Println(f(), f(), f())

	s, _ := Extract("google.com")
	fmt.Println(s)
}
