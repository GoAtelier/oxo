package main

import (
	"fmt"

	"github.com/finecoding/oxopfork/oxo"
)

func main() {

	// This is a quick test to see if the internal import is working
	// oxo package contains types.go which defines Grid as [9]byte

	// Added a stringer function each for Grid, Turn and Game to assist printing
	// Put them new file stringer.go in oxo package.  Copied much of it from oxo2.go

	b := oxo.Grid([]byte("X........"))
	fmt.Println("This is a board", b)
	t := oxo.Turn{
		Board:  b,
		Status: "XWIN"}
	fmt.Println("This is a turn", t)
	g := oxo.Game{t, t, t, t, t}
	fmt.Println("This is a game", g)
}
