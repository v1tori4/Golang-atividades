// - Crie um map com key tipo string e value tipo []string.
//     - Key deve conter nomes no formato sobrenome_nome
//     - Value deve conter os hobbies favoritos da pessoa
// - Demonstre todos esses valores e seus índices.

package main

import (
	"fmt"
)

//main
func hh() {

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

	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}

}

//resultado:
//strangerthings_eleven
// 0  -  desaparecer
// 1  -  ser assustadora
// 2  -  raspar o cabelo
// senna_ayrton
// 0  -  andar de jetski
// 1  -  ser milionário
// 2  -  falar com sotaque de paulista mano
// pike_rob
// 0  -  criar linguagens de programação
// 1  -  usar uns ternos muito loucos
