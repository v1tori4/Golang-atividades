//slice: Maps: introdução
//- Maps *não tem ordem.*
//****

package main

import "fmt"

//main
func jjj() {
	amigos := map[string]int{
		"alfredo": 5551234,
		"joana":   9996674,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["joana"])

	amigos["gopher"] = 444444

	fmt.Println(amigos)
	fmt.Println(amigos["gopher"], "\n\n")

	// comma ok idiom
	if será, ok := amigos["fantasma"]; !ok {
		fmt.Println("não tem!")
	} else {
		fmt.Println(será)
	}
}

//resultado:

// map[alfredo:5551234 joana:9996674]
// 9996674
// map[alfredo:5551234 gopher:444444 joana:9996674]
// 444444

// não tem!
