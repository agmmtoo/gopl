package main

import (
	"fmt"
	"github.com/agmyintmyatoo/gopl/ch6"
)

func main() {
	var x, y ch6.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWidth(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123))
}
