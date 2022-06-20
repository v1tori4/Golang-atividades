//slice: fatiando ou deletando de uma fatia
//****
//desafio: mostrar todos os sabores sem usar o range

package main

import "fmt"

//main
func dd() {
	//                  0.            1.          2.         3.                4.
	sabores := []string{"pepperoni", "mussarela", "abacaxi", "quatro queijos", "marguerita"}
	//jeito 1:
	fatia := sabores[:]
	fmt.Println(fatia)

	//jeito2
	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}

}
