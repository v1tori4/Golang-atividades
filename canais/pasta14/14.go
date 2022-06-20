package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	ímpar := make(chan int)
	converge := make(chan int)

	go envia(par, ímpar)
	go recebe(par, ímpar, converge)

	for v := range converge {
		fmt.Println("Valor recebido:", v)
	}

}

func envia(p, i chan int) {
	x := 100
	for n := 0; n < x; n++ {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	close(p)
	close(i)
}

func recebe(p, i, c chan int) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for v := range p {
			c <- v
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for v := range i {
			c <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(c)
}

// Valor recebido: 0
// Valor recebido: 2
// Valor recebido: 1
// Valor recebido: 3
// Valor recebido: 4
// Valor recebido: 5
// Valor recebido: 6
// Valor recebido: 7
// Valor recebido: 8
// Valor recebido: 9
// Valor recebido: 10
// Valor recebido: 11
// Valor recebido: 12
// Valor recebido: 13
// Valor recebido: 14
// Valor recebido: 15
// Valor recebido: 16
// Valor recebido: 17
// Valor recebido: 18
// Valor recebido: 19
// Valor recebido: 20
// Valor recebido: 21
// Valor recebido: 22
// Valor recebido: 23
// Valor recebido: 24
// Valor recebido: 25
// Valor recebido: 26
// Valor recebido: 27
// Valor recebido: 28
// Valor recebido: 29
// Valor recebido: 30
// Valor recebido: 31
// Valor recebido: 32
// Valor recebido: 33
// Valor recebido: 34
// Valor recebido: 35
// Valor recebido: 36
// Valor recebido: 37
// Valor recebido: 38
// Valor recebido: 39
// Valor recebido: 40
// Valor recebido: 41
// Valor recebido: 42
// Valor recebido: 43
// Valor recebido: 44
// Valor recebido: 45
// Valor recebido: 46
// Valor recebido: 47
// Valor recebido: 48
// Valor recebido: 49
// Valor recebido: 50
// Valor recebido: 51
// Valor recebido: 52
// Valor recebido: 53
// Valor recebido: 54
// Valor recebido: 55
// Valor recebido: 56
// Valor recebido: 57
// Valor recebido: 58
// Valor recebido: 59
// Valor recebido: 60
// Valor recebido: 61
// Valor recebido: 62
// Valor recebido: 63
// Valor recebido: 64
// Valor recebido: 65
// Valor recebido: 66
// Valor recebido: 67
// Valor recebido: 68
// Valor recebido: 69
// Valor recebido: 70
// Valor recebido: 71
// Valor recebido: 72
// Valor recebido: 73
// Valor recebido: 74
// Valor recebido: 75
// Valor recebido: 76
// Valor recebido: 77
// Valor recebido: 78
// Valor recebido: 79
// Valor recebido: 80
// Valor recebido: 81
// Valor recebido: 82
// Valor recebido: 83
// Valor recebido: 84
// Valor recebido: 85
// Valor recebido: 86
// Valor recebido: 87
// Valor recebido: 88
// Valor recebido: 89
// Valor recebido: 90
// Valor recebido: 91
// Valor recebido: 92
// Valor recebido: 93
// Valor recebido: 94
// Valor recebido: 95
// Valor recebido: 96
// Valor recebido: 97
// Valor recebido: 98
// Valor recebido: 99
