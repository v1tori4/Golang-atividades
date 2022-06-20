// Crie e use um struct anônimo. OK
// Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.

package main

import "fmt"

//main
func dd() {
	pessoa := struct {
		nome         string
		sobrenome    string
		gostadecomer []string
		gostadebeber map[string]string
	}{
		nome:         "Vitória",
		sobrenome:    "Nascimento",
		gostadecomer: []string{"Pamonha", "Cural"},
		gostadebeber: map[string]string{
			"No cafe da manhã": "suco de milho",
			"No almoço":        "suco de goiaba",
		},
	}
	fmt.Println(pessoa)
}
