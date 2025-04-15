package main

import (
	"fmt"
	"math/rand"
)

func ReturnHello() string {
	return "Hello, World!"
}

func ReturnInteger() int {
	return rand.Int()
}

func main() {
	fmt.Println(ReturnHello())
	fmt.Printf("Random Integer: %d\n", ReturnInteger())
}
