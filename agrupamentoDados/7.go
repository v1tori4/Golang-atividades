//slice: slice multi-dimencional
//agrupamento de dados
//****

package main

import "fmt"

//main
func hhh() {
	ss := [][]int{
		// Índice: 0  1  2                   // Índice:
		[]int{1, 2, 3, 4, 5, 6},       // 0
		[]int{7, 8, 9, 10, 11, 12},    // 1
		[]int{13, 14, 15, 16, 17, 18}, // 2
	}
	fmt.Println(ss[2][4]) //pego a subfatia numero 2 e o indice numero 4
}
