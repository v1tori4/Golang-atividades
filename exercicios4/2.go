// Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// Demonstre os valores do map utilizando range.
// Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

//main
func main() {

	meumapa := make(map[string]pessoa)

	meumapa["Nascimento"] = pessoa{
		nome:      "Vitória",
		sobrenome: "Nascimento",
		sabores:   []string{"Chocolate", "Baunilha", "Maracuja", "Uva"}}

	meumapa["Pereira"] = pessoa{"Joana", "Pereira",
		[]string{"Flocos", "Abacaxi", "Pistache", "Morango"}}

	for _, valor := range meumapa {
		fmt.Println("Meu nome é", valor.nome, valor.sobrenome, "e meus sorvetes favoritos são:")
		for _, valor := range valor.sabores {
			fmt.Println("–", valor)
		}
	}

}

// resultado:
// Meu nome é Vitória Nascimento e meus sorvetes favoritos são:
// – Chocolate
// – Baunilha
// – Maracuja
// – Uva
// Meu nome é Joana Pereira e meus sorvetes favoritos são:
// – Flocos
// – Abacaxi
// – Pistache
// – Morango
