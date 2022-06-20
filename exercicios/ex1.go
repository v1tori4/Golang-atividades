package main

import "fmt"

// %d- decimal,
// %#x- hexadecimal,
// %b- binario,
// %s- string

func main() {
	x := 100
	fmt.Printf("%d, %#x, %b", x, x, x)
}
