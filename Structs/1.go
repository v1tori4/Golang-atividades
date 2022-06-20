// Struct é um tipo de dados composto que nos permite armazenar valores de tipos *diferentes.*
// - Declaração: type x struct { y: z }
// - Acesso: x.y
// - Exemplo: nome, idade, fumante.

package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

//main
func aa() {
	cliente1 := cliente{
		nome:      "João",
		sobrenome: "da Silva",
		fumante:   false,
	}

	cliente2 := cliente{"Joana", "Pereira", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)

}
