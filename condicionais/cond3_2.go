//default: se nada acontecer, executa ele
//varias comparações

package main

import "fmt"

//main
func cc() {
	quemtanoescritoriohoje := "ninguem"

	switch quemtanoescritoriohoje {
	case "zezinho", "maria":
		fmt.Println("hoje quem ta na firma é o time 1")
	case "joana", "pedrinho":
		fmt.Println("hoje quem ta na firma é o time 2")
	default:
		fmt.Println("ta vazio")
	}
}
