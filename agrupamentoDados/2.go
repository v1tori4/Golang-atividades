//slice: literal composta
// slice é um array
//conjunto de dados que podem ser int, float, string...
//o array tem um numero pré definido e o slice não

package main

import "fmt"

//main
func bbb() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2 := append(slice, 6)
	fmt.Println(slice2)

	fmt.Println(slice[3]) //pegou o valor 4 do indice 3
	slice[3] = 348928     //atribuiu outro valor
	fmt.Println(slice[3])
}
