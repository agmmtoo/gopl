package main

import (
	"fmt"
	"bytes"
	//"strconv"
)

func comma(n int) string {
	num := string(n)
	var buf bytes.Buffer
	for i := len(num)-1; i >= 0; i-- {
		if len(buf.String()) % 3 == 0 {
			buf.WriteString(", ")
			fmt.Println(len(buf.String()))

		}
		buf.WriteByte(num[i])
		fmt.Println(num[i])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma(123456), comma(12), comma(123), comma(12345))
}
