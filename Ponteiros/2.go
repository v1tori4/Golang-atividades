// Ponteiros – 2. Quando usar ponteiros
// Ponteiros permitem compartilhar endereços de memória. Isso é útil quando:
//     - Não queremos passar grandes volumes de dados pra lá e pra cá
//     - Queremos mudar um valor em sua localização original
//***

package main

import "fmt"

func main() {
	x := 11

	//estarecebeovalor(x)
	estarecebeumponteiro(&x)

	fmt.Println(x)

}

func estarecebeovalor(x int) {
	x++
	fmt.Println("Na função:", x)

}

func estarecebeumponteiro(x *int) {
	*x++
	fmt.Println("Na função:", *x)
}
