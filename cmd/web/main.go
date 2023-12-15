package main

import (
	"fmt"

	"github.com/finecoding/oxo-1/oxo"
)

func main() {

	board := oxo.Grid([]byte("X........"))
	fmt.Println(board)
}
