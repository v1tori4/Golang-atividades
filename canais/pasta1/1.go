package main

import "fmt"

func main() {
	canal := make(chan int)
	canal <- 42
	fmt.Println(<-canal)
}
