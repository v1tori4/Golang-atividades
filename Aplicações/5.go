//Sort com string

package main

import (
	"fmt"
	"sort"
)

//main
func d() {

	ss := []string{"abóbora", "maçã", "laranja", "beringela", "berinjela"}

	fmt.Println(ss)

	sort.Strings(ss)

	fmt.Println(ss)

}
