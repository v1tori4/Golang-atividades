// exercicio 9: callback

package main

import "fmt"

//main
func ii() {
	essavaireceberaoutrafuncao(issovaiserumargumento)
}

func issovaiserumargumento() {
	fmt.Println("Olha eu aqui!")
}

func essavaireceberaoutrafuncao(x func()) {
	fmt.Println("Prestenção:")
	x()
}
