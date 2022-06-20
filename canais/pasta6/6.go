package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) //recebe
	cs := make(chan<- int) //envia

	fmt.Println(".......")
	fmt.Println("c\t%T\n", cr)
	fmt.Println("c\t%T\n", cs)

	//geral para conversão especificas
	fmt.Println(".....")
	fmt.Println("c\t%T\n", (<-chan int)(c))
	fmt.Println("c\t%T\n", (chan<- int)(c))
}
