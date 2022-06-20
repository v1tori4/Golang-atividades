// - Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos: OK
//Nome OK
//Sobrenome OK
//Sabores favoritos de sorvete OK
//Crie dois valores do tipo "pessoa" e demonstre estes valores, OK
//utilizando range na slice que contem os sabores de sorvete.OK

package main

import "fmt"

// type pessoa struct {
// 	nome      string
// 	sobrenome string
// 	sabores   []string
// }

//main
func aa() {

	pessoa1 := pessoa{
		nome:      "Vitória",
		sobrenome: "Nascimento",
		sabores:   []string{"Chocolate", "Baunilha", "Maracuja", "Uva"},
	}

	pessoa2 := pessoa{"Joana", "Pereira", []string{"Flocos", "Abacaxi", "Pistache", "Morango"}}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, s := range pessoa1.sabores {
		fmt.Println("\t", s)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, s := range pessoa2.sabores {
		fmt.Println("\t", s)
	}

}

// resultado:
// Vitória Nascimento
//          Chocolate
//          Baunilha
//          Maracuja
//          Uva
// Joana Pereira
//          Flocos
//          Abacaxi
//          Pistache
//          Morango
