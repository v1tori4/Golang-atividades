//Struct: embutidos
//Structs dentro de structs dentro de structs.

package main

import "fmt"

// type pessoa struct {
// 	nome  string
// 	idade int
// }

// type profissional struct {
// 	pessoa
// 	titulo  string
// 	salario int
// }

//main
func bb() {

	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}
	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Marcio",
			idade: 62,
		},
		titulo:  "Professor",
		salario: 10000,
	}
	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
}

// resultado:
// {Alfredo 30}
// {{Marcio 62} Professor 10000}
