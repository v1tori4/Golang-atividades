// - Começando com a seguinte slice:
//     - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// - Anexe a ela o valor 52;
// - Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
// - Demonstre a slice;
// - Anexe a ela a seguinte slice:
//     - y := []int{56, 57, 58, 59, 60}
// - Demonstre a slice x.

//????????

package main

import "fmt"

//main
func dd() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x[:3], x[len(x)-4:]...) //42, 43, 44, 45, 46, 47
	x = append(x[:3], x[6:]...)

	fmt.Println(x)
}
