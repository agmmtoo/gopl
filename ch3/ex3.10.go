package main

import (
	"bytes"
	"fmt"
)

func comma(values []int) string {
	var buf bytes.Buffer
	for i, v := range values {
		if (len(values))-i > 0 && (len(values)-i)%3 == 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	return buf.String()
}

func main() {
	fmt.Println(comma([]int{2, 3, 4, 5, 6, 7, 8, 9}))
}
