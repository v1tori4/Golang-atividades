package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)

	go func() {
		canal <- 0
		close(canal)
	}()

	v, ok := <-canal

	fmt.Println(v, ok)

	v, ok = <-canal

	fmt.Println(v, ok)
}

//0 true
//0 false
