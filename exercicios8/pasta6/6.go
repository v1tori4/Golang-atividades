package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Seu OS é:\t", runtime.GOOS)
	fmt.Println("Seu ARCH é:\t", runtime.GOARCH)
}

// Seu OS é:        windows
// Seu ARCH é:      amd64
