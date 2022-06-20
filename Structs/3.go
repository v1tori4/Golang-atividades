//Struct: lendo a documentação

package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

//main
func dd() {

	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Maricota",
			idade: 31,
		},
		titulo:  "Pizzaiola",
		salario: 10000,
	}

	pessoa3 := pessoa{"Mauricio", 40}
	pessoa4 := profissional{pessoa{"Vanderlei", 50}, "Político", 10000000}

	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.pessoa.nome)
	fmt.Println(pessoa3.nome)
	fmt.Println(pessoa4)
}

// resultado:
// Alfredo
// Maricota
// Mauricio
// {{Vanderlei 50} Político 10000000}
