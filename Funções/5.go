// Funções – 5. Interfaces & polimorfismo
// Em Go, valores podem ter mais que um tipo.
// Declaração: keyword identifier type → type x interface
// - Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.
// - Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum)
// então esse tipo implicitamente implementa a interface.
// - Esse tipo será o seu tipo *e também* o tipo da interface.

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesarrancados int
	salário          float64
}

type arquiteto struct {
	pessoa
	tipodeconstrução string
	tamanhodaloucura string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e eu já arranquei", x.dentesarrancados, "dentes, e ouve só: Bom dia!")
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e ouve só: Bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
	switch g.(type) {
	case dentista:
		fmt.Println("Eu ganho:", g.(dentista).salário)

	case arquiteto:
		fmt.Println("Eu construo:", g.(arquiteto).tipodeconstrução)
	}
}

//main
func ee() {
	mrdente := dentista{
		pessoa: pessoa{
			nome:      "Mister Dente",
			sobrenome: "da Silva",
			idade:     50,
		},
		dentesarrancados: 8973,
		salário:          98736.06,
	}

	mrprédio := arquiteto{
		pessoa: pessoa{
			nome:      "Mister Prédio",
			sobrenome: "Sobrenome",
			idade:     51,
		},
		tipodeconstrução: "Hospícios",
		tamanhodaloucura: "Demais",
	}

	serhumano(mrdente)
	fmt.Println("")
	serhumano(mrprédio)

}
