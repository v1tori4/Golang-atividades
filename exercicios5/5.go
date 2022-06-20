//exercicio 5:
// - Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
// - Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
// - Crie um valor de tipo "quadrado"
// - Crie um valor de tipo "círculo"
// - Use a função "info" para demonstrar a área do "quadrado"
// - Use a função "info" para demonstrar a área do "círculo"

package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

func (q quadrado) area() {
	resultado := q.lado * q.lado
	fmt.Println("Área do quadrado:", resultado)
}

type circulo struct {
	raio float64
}

func (c circulo) area() {
	resultado := math.Pi * 2 * c.raio
	fmt.Println("Área do círculo:", resultado)
}

type triangulo struct {
	base   float64
	altura float64
}

func (c triangulo) area() {
	resultado := c.base * c.altura / 2
	fmt.Println("Área do triangulo:", resultado)
}

type info interface {
	area()
}

func medida(i info) {
	i.area()
}

//main
func main() {

	x := quadrado{
		lado: 15.0,
	}

	y := circulo{
		raio: 5.25,
	}

	z := triangulo{
		base:   10,
		altura: 9,
	}

	medida(x)

	medida(y)

	medida(z)
}
