//exercicio 1:

package main

import "fmt"

//main
func aa() {
	fmt.Println(retornaumint())
	fmt.Println(retornauminteumastring())
}

func retornaumint() int {
	return 10
}

func retornauminteumastring() (int, string) {
	return 20, "vinte"
}
