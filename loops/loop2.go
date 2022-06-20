//nested loop (repetição hierárquica)**
//ciclos pequenos em ciclos grandes

package main

import "fmt"

func main() {
	for mes := 1; mes <= 12; mes++ {
		fmt.Println("Mes: ", mes)

		for x := 1; x < 30; x++ {
			fmt.Println("Dia: ", x, ", ")
		}
		fmt.Printf("\n")
	}

}
