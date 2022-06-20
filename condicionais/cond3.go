//Condicionais: a declaração switch
//é como se fosse um "if" com varios focos diferentes
//fallthrough : sempre quando é executado, pula a comparação e executa o próximo

package main

import "fmt"

func main() {
	x := 2

	switch true {
	case x < 5:
		fmt.Println("x é menor que 5")
		fallthrough
	case x == 5:
		fmt.Println("x é igual a 5")
	case x > 5:
		fmt.Println("x é maior que 5")

	}
}
