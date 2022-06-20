//loops: break e continue
//modulo: resto da divisao (10 modulo 3, vou pegar 3 dividir por 10 e sobra 1)
//o BREAKE quebra o loop como um todo, e o CONTINUE quebra aquel interação especifica do loop

package main

import "fmt"

func main() {

	//verificar numeros que são divisiveis por 2, é par

	for i := 0; i < 20; i++ {
		if i == 3 {
			//faz isso quando o numero nao é par, paro a interação do loop e vou direto para a proxima
			continue
		}
		fmt.Println(i)
	}
}
