//criar um loop utilizando a sintaxe: for condition
//mostrar os anos desde que nasceu

package main

import "fmt"

//main
func loop() {
	ano := 20003
	anofinal := 20103
	for ano <= anofinal {
		fmt.Println(ano)
		ano++
	}

}
