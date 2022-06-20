// exercicios 3:

package main

import "fmt"

//main
func cc() {
	defer fmt.Println("isso vai aparecer depois")
	fmt.Println("isso vai aparecer antes")
}
