// Funções – 8. Retornando uma função
// Declaração: func f() return
// Exemplo: func f() func() int { [...]; return func() int{ return [int] } }
// ????: fmt.Println(f()())

package main

import "fmt"

//main
func ii() {

	x := retornaumafuncao()

	y := x(3)

	fmt.Println(y)

}

func retornaumafuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
