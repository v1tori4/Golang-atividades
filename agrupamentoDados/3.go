//slice: for range
//ela atravessa toda a extensao do slice e mostra oq tem la
//append adiciona slice
//****

package main

import "fmt"

//main
func cc() {
	slice := []string{"banana", "maçã", "jaca", "pessego"}
	for indice, valor := range slice {
		fmt.Println("No indice", indice, "temos o valor:", valor)
	}

	slice = append(slice, "melancia")
	for _, valor := range slice {
		fmt.Println("Um dos valores desse sile é:", valor)
	}

}
