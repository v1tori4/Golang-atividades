//Condicionais: a declaração switch pt 2. e documentação

package main

import "fmt"

var x interface{}

//main
func declaracao() {
	x = "batman"
	switch x.(type) {
	case int:
		fmt.Println("é um int")
	case bool:
		fmt.Println("é um boolean")
	case string:
		fmt.Println("é uma string")
	case float64:
		fmt.Println("é um float")
	}
}
