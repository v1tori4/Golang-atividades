package main

import "fmt"

func main() {
	a := 100
	fmt.Printf("%d\t %#x\t %b\t \n", a, a, a)
	//

	b := a << 1
	fmt.Printf("%d\t %#x\t %b\t", b, b, b)
	//deslocar bits da variavel 1 para a esquerda
	//atribuir o resultado para outra variavel
	// ">>" multiplica/ "<<" divisao
}
