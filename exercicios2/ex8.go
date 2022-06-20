//ultilizar a declaração switch sem switch statement especificado (qualquer coisa que for true ta valendo)

package main

import "fmt"

//main
func aaa() {
	tocomfome := 2

	switch {

	case tocomfome == 0:
		fmt.Println("Ainda nao comi")
	case tocomfome == 1:
		fmt.Println("Ja comi")
	case tocomfome == 2:
		fmt.Println("Quero sobremesa")

	}

}
