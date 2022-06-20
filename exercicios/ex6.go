package main

import "fmt"

// o que é iota

const (
	_ = 2003 + iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(b, c, d, e)
}
