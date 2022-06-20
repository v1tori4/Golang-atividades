//a delcaração for
//loop while não existe, apenas for
//se retirar o "x++", o servidor roda pra sempre
//se quiser um loop infinito, tirar a condição ( x < 10)

package main

import "fmt"

func main() {
	x := 0

	for x < 10 {
		fmt.Println("x é menor que 10")
		x++
	}

}
