//slice: a surpresa do array subjacente
//Toda slice tem um array subjacente.
//Um slice é: um ponteiro/endereço para um array, mais len e cap (que é o len to array).
//****

package main

import "fmt"

//main
func jj() {
	primeiroslice := []int{1, 2, 3, 4, 5}

	fmt.Println(primeiroslice)

	segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)

	fmt.Println(segundoslice)

	fmt.Println(primeiroslice)
}

//resultado:

//[1 2 3 4 5]
//[1 2 5]
//[1 2 5 4 5]
