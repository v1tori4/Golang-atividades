// - Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional
//da seguinte maneira:
//     - "Nome", "Sobrenome", "Hobby favorito"
// - Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

package main

import "fmt"

//main
func ff() {
	ss := [][]string{
		[]string{
			"Paulo",
			"Batista",
			"Cantar",
		},
		[]string{
			"Martha",
			"Silva",
			"Fazer comida",
		},
		[]string{
			"Geysa",
			"Amaro",
			"Costurar",
		},
	}

	//mostrando os slices que tem dentro do slice
	for _, v := range ss {
		fmt.Println(v)
	}

	fmt.Println("\n\n")

	//mostrando cada item que tem no slice
	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}

	}

}
