// exercicio 4:

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (funcionario pessoa) demonstrar() {
	fmt.Println("O nome completo dessa pessoa é:", funcionario.nome, funcionario.sobrenome,
		"\nEssa pessoa tem", funcionario.idade, "anos.")
}

//main
func dd() {

	funcionario2 := pessoa{
		nome:      "Maria",
		sobrenome: "Lurdes",
		idade:     46,
	}

	funcionario2.demonstrar()

}
