// exercicio 10: closure
///***

package main

import "fmt"

//main
func jj() {
	a := i()
	b := i()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

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
