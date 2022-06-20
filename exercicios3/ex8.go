// Utilizando o exercício anterior,
//adicione uma entrada ao map e demonstre o map inteiro utilizando range.

package main

import "fmt"

func main() {

	mepe := map[string][]string{
		"strangerthings_eleven": []string{
			"desaparecer", "ser assustadora", "raspar o cabelo",
		},
		"senna_ayrton": []string{
			"andar de jetski", "ser milionário", "falar com sotaque de paulista mano",
		},
		"pike_rob": []string{
			"criar linguagens de programação", "usar uns ternos muito loucos",
		},
	}

	mepe["loureiro_kiko"] = []string{"usar os trequinho no punho", "tacar fogo na guitarra"}

	delete(mepe, "strangerthings_eleven")

	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}

}
