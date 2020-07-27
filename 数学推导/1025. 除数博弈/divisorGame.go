package main

import (
	"fmt"
)

func main() {
	N := 105
	fmt.Println(divisorGame(N))
}

func divisorGame(N int) bool {
	return N%2 == 0
}
