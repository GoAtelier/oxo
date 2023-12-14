package main

import (
	"fmt"

	"github.com/finecoding/oxopfork/oxo"
)

func main() {

	board := oxo.Grid([]byte("X........"))
	fmt.Println(board)
}
