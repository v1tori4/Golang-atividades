// Funções – 9. Callback
// Exemplo:
// Criando uma função que toma uma função e um []int, e usa somente os números pares como argumentos para a função.

package main

import "fmt"

//main
func jj() {
	t := somentePares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println(t)
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func somentePares(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if (v % 2) == 0 {
			slice = append(slice, v)
		}
		total := f(slice...)
		return total
		// if v%2 != 0(impar)
	}
}

//faz isso aqui

// total := f(slice...)
// return total

// func somentePares(f func(x ...int) int, y ...int) int {

// 	quando isso aqui acabar:

// 	var slice []int
// for _, v := range y {
// 	if v%2 == 0 {
// 		slice = append(slice, v)
// 	}
// }
