//slice: anexando a uma slice
//entendendo append
//****
//varidiac

package main

import "fmt"

//main
func ff() {
	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12}

	fmt.Println(umaslice)

	umaslice = append(umaslice, 5, 6, 7, 8)

	fmt.Println(umaslice)

	umaslice = append(umaslice, outraslice...) //"..." pagr os itens que estão dentro da "outraslice"

	fmt.Println(umaslice)

}
