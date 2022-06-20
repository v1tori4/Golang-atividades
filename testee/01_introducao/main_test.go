package main

import (
	"testing"
)

//esse da certo
func TestSoma(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

//esse falha
// func TestSoma2(t *testing.T) {
// 	teste := soma(3, 2, 1)
// 	resultado := 7
// 	if teste != resultado {
// 		t.Error("Expected:", resultado, "Got:", teste)
// 	}
// }

func TestMultiplica(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}
