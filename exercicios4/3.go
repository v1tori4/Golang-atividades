// Crie um novo tipo: veículo OK
// - O tipo subjacente deve ser struct OK
// - Deve conter os campos: portas, cor OK
// - Crie dois novos tipos: caminhonete e sedan OK
// - Os tipos subjacentes devem ser struct OK
// - Ambos devem conter "veículo" como struct embutido OK
// - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro" OK
// - O tipo sedan deve conter um campo bool chamado "modeloLuxo" OK
// - Usando os structs veículo, caminhonete e sedan: OK
// - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos OK
// - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos OK
// - Demonstre estes valores. OK
// - Demonstre um único campo de cada um dos dois. OK

package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	traçãoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

//main
func cc() {

	veiculo1 := veiculo{

		portas: 4,
		cor:    "Branco",
	}

	veiculo2 := caminhonete{
		veiculo: veiculo{
			portas: 2,
			cor:    "Preto",
		},
		traçãoNasQuatro: true,
	}

	veiculo3 := sedan{
		veiculo: veiculo{
			portas: 4,
			cor:    "Vermelho",
		},
		modeloLuxo: true,
	}

	fmt.Println(veiculo1)
	fmt.Println(veiculo2)
	fmt.Println(veiculo3)

}

// {solução dela
// type veículo struct {
// 	portas int
// 	cor    string
// }

// type caminhonete struct {
// 	veículo
// 	traçãoNasQuatro bool
// }

// type sedan struct {
// 	veículo
// 	modeloLuxo bool
// }

// func main() {
// 	carrãodotio := sedan{veículo{4, "abóbora"}, true}
// 	fubicadovô := caminhonete{
// 		veículo: veículo{
// 			portas: 8,
// 			cor:    "ferrugem",
// 		},
// 		traçãoNasQuatro: false,
// 	}

// 	fmt.Println(carrãodotio)
// 	fmt.Println(fubicadovô)
// 	fmt.Println(carrãodotio.cor)
// 	fmt.Println(fubicadovô.traçãoNasQuatro)

// }
// }
