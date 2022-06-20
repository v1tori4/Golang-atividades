// Partindo do código abaixo, utilize marshal para transformar []user em JSON.

// type user struct {
// 	first string
// 	age   int
// }

// func main() {
// 	u1 := user{
// 		first: "James",
// 		age:   32,
// 	}

// 	u2 := user{
// 		first: "Moneypenny",
// 		age:   27,
// 	}

// 	u3 := user{
// 		first: "M",
// 		age:   54,
// 	}

// 	users := []user{u1, u2, u3}

// 	fmt.Println(users)

// }

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

//main
func a() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	sb, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

}
