// Funções – 10. Closure

package main

import "fmt"

//main
func kk() {

	a := i()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := i()

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
