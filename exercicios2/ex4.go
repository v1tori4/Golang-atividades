//criar um loop infinito

package main

import "fmt"

//main
func infinito() {
	ano := 20003
	anofinal := 20103
	for {
		if ano > anofinal {
			break
		}
		fmt.Println(ano)
		ano++
	}
}
