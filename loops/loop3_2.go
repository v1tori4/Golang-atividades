//o "break" quebra o loop

package main

import "fmt"

func main() {
	x := 0

	for {
		if x < 10 {
			x++
			fmt.Println("x é menor que 10")
		} else {
			fmt.Println("x é maior que 10")
			break
		}

	}
	fmt.Println("o loop esta brekado!")
}
