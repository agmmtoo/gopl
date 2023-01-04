package main

import (
	"fmt"
	"buffer"
)

func comma(num int) string {
	num = string(num)
	var buf buffer.Buffer
	for i := len(num); i > 0; i-- {
		if len(buf) % 3 == 0 {
			buf.WriteString(", ")
		}
		buf.WriteInt(num[i])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma(123456), comma(12), comma(123), comma(12345))
}
