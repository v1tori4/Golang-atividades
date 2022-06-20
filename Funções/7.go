// Funções – 7. Func como expressão
// f := func(p params){ ... }
// - f()package main

package main

import "fmt"

//main
func hh() {

	x := 2

	y := func(x int) int {
		//fmt.Println(x, "vezes 4 é:")
		return x * 4
	}

	fmt.Println(x, "vezes 4 é:", y(x))

}
