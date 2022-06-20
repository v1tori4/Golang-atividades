package main

import "fmt"

func main() {
	a := (10 == 10)
	b := (10 != 100)
	c := (10 <= 100)
	d := (10 < 100)
	e := (100 >= 10)
	f := (100 > 10)

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)

}
