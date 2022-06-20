// exercicio 8:

package main

import "fmt"

//main
func hh() {

	x := retornaumafunc()

	x()
}

func retornaumafunc() func() {
	return func() {
		fmt.Println("Olha eu aqui!")
	}
}
