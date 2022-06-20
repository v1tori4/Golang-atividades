//Condicionais: if, else if, else

package main

import "fmt"

//main
func bb() {
	if x := 500; x > 100 {
		fmt.Println("X é maior que 100")
	} else if x < 10 {
		fmt.Println("X é menor que 10")
	} else {
		fmt.Println("X não é menor que 10 nem maior que 100")
	}
}
