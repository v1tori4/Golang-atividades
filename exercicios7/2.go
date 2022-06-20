//utilize unmarshal e demonstre os valores.

package main

import (
	"encoding/json"
	"fmt"
)

type jsontogo []struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

//main
func b() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Hi", "How are you?"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["I´m good",", and you?"]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["I´m fine :)"]}]`
	fmt.Println(s)

	var resultado jsontogo

	err := json.Unmarshal([]byte(s), &resultado)
	if err != nil {
		fmt.Println("erro", err)
	}

	fmt.Println(resultado)
	fmt.Println(resultado[1])
	fmt.Println(resultado[1].Last)

}
