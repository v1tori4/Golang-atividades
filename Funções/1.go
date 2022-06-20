//funcoes: 1. Sintaxe
//Função básica
//Função que aceita um argumento
//Função com retorno(soma simples)
//Função com múltiplos retornos e parâmetro variádico. (função que retorna o total de todos os ints recebidos)
//func (receiver) nome (parametro) return {code}

package main

import "fmt"

//main
func aa() {
	basica()

	argumento("manhã")
	argumento("tarde")
	argumento("noite")

	valor := soma(10, 10)
	fmt.Println(valor)

	total, quantos := conta(1, 5, 4, 3)
	fmt.Println(total, quantos)

}

func basica() {
	fmt.Println("Seja bem vindo!")
}

func argumento(s string) {
	if s == "manhã" {
		fmt.Println("Oi, bom dia!")
	} else if s == "tarde" {
		fmt.Println("Oi, boa tarde!")
	} else {
		fmt.Println("Oi, boa noite!")
	}
}

func soma(x, y int) int {
	return x + y
}

func conta(x ...int) (int, int) {
	conta := 0
	for _, v := range x {
		conta += v
	}
	return conta, len(x) //retorna o valor da soma e a quantidade de elementos
}
