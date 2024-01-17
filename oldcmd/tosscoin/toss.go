package main

import (
	"fmt"
	//"math/rand"
)

func TossCoin(t bool) bool {
//
//	if flipint := rand.Intn(2); flipint == 0 {
//		return true
//	}
//	return false
return !t
}
func main() {
	t:=true
var o,x int
	for turn := 0; turn < 1000; turn++ {
		t=!t
		if TossCoin(t) {
			o++
			fmt.Printf("O")

		} else {
			x++
			fmt.Printf("X")
		}
	}
	fmt.Printf("\nX = %d \nO = %d ",x,o)
}
