// Funções – 11. Recursividade
// - Exemplo: fatoriais.
// Com recursividade:

package main

import "fmt"

//main
func ll() {
	fmt.Println(funcao(6))
}

func funcao(x int) int {
	if x == 1 {
		return x
	}
	return x * funcao(x-1)
}
