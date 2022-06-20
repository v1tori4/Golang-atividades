//slice: fatiando ou deletando de uma fatia (pt. 2)
//deletando o sabor "abacaxi"
//****
//

package main

import "fmt"

//main
func ee() {
	//                  0.            1.          2.         3.                4.
	sabores := []string{"pepperoni", "mussarela", "abacaxi", "quatro queijos", "marguerita"}

	sabores = append(sabores[:2], sabores[3:]...)
	fmt.Println(sabores)

}
