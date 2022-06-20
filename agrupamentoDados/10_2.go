//Maps: range e deletando
//Reiterando: maps *não tem ordem* e um range usará uma ordem aleatória.
// delete(map, key)
// Deletar uma key não-existente não retorna erros!

package main

import "fmt"

//main
func gmain() {

	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		19:  "idade de ir pra festa",
	}

	fmt.Println(qualquercoisa)

	total := 0

	for key, _ := range qualquercoisa {
		total += key
	}

	fmt.Println(total)

	delete(qualquercoisa, 123)

	fmt.Println(qualquercoisa)

}
