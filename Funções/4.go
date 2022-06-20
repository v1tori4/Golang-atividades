// Funções – 4. Métodos
// - Quando se anexa uma função a um tipo, ela se torna um método desse tipo.
// - Pode-se anexar uma função a um tipo utilizando seu receiver.
// - Utilização: valor.método()

package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

//funcao especifica para valores daquele tipo
func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "diz bom dia!")
}

//main
func dd() {

	mauricio := pessoa{"Maurício", 30}
	mauricio.oibomdia() //funcao anexada a um tipo
}

// func (receiver) identifier(parameters) (returns) { code }
