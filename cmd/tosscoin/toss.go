package main

import (
	"fmt"
	"math/rand"
)

func TossCoin() bool {

	if flipint := rand.Intn(2); flipint == 0 {
		return true
	}
	return false
}
func main() {
var o,x int
	for turn := 0; turn < 1000; turn++ {
		if TossCoin() {
			o++
			fmt.Printf("O")

		} else {
			x++
			fmt.Printf("X")
		}
	}
	fmt.Printf("\nX = %d \nO = %d ",x,o)
}
