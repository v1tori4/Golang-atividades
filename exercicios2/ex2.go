//mostrar letra A ate Z, 3 vezes

package main

import "fmt"

//main
func mostrar() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U", i)
		}
	}

}
