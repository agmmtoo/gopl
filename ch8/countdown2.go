package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Comencing count")
	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		case <-tick:
			// Do nothing.
		}
	}
}
