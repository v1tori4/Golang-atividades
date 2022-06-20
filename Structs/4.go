// Struct: anonimo
// não declara o tipo e não pode ser usado varias vezes

package main

import "fmt"

//main
func ee() {

	//nao criou um tipo, apenas definiu uma estrutura e ultilizou
	x := struct {
		nome  string
		idade int
	}{
		nome:  "Mario",
		idade: 50}

	fmt.Println(x)
}
