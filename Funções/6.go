// Funções – 6. Funções anônimas
// Anonymous self-executing functions → Funções anônimas auto-executáveis.
// func(p params) { ... }()
// Vamos ver bastante quando falarmos de goroutines.

package main

import "fmt"

//main
func gg() {

	x := 2

	func(x int) {
		fmt.Println(x, "vezes 6 é:")
		fmt.Println(x * 6)
	}(x)

}
