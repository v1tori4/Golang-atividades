package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println(p.nome, "diz oi!")
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	p1 := pessoa{"João Silva", 10}

	p1.falar()

	dizerAlgumaCoisa(&p1)
}

// João Silva diz oi!
// João Silva diz oi!
