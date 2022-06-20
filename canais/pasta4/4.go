package main

import "fmt"

func main() {
	canal := make(chan int, 1)
	canal <- 42
	canal <- 43
	fmt.Println(<-canal)
}
