//Funções – 2. Desenrolando (enumerando) uma slice
//variadica: pode ter uma quantidade militada ou 0
//Quando temos uma slice, podemos passar os elementos individuais através "deste..." operador

package main

import "fmt"

//main
func bb() {

	slice := []int{10, 10, 1, 2, 3, 5}

	total := adicao(slice...)

	fmt.Println(total)
}

func adicao(x ...int) int {
	adicao := 0
	for _, v := range x {
		adicao += v
	}
	return adicao
}
