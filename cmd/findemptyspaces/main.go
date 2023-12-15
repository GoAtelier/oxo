package main

import (
	"fmt"

	"github.com/finecoding/oxopfork/oxo"
)

// Note this is using a grid as a byte slice, so it is old.
// The Inplay function returns a list of empty spaces on a 3x3 grid.
// The Grid type was changed to an array
// It needs to a rewrite and maybe add a function to choose and empty space at random
// This will be the basis of the simplest tactic.
// We need two players, each with a tactic to play a game.

func main() {

	var board oxo.Grid

	board = oxo.Grid([]byte("OO XXX OO"))
	fmt.Printf("%s", board)
	a, b := board.InPlay()
	fmt.Printf("%v, %v", a, b)

	board = oxo.Grid([]byte("OO0XXXXOO"))
	fmt.Printf("%s", board)
	a, b = board.InPlay()

	fmt.Printf("%v, %v", a, b)
}

// InPlay takes a board byte slice and returns a byte slice of empty positions
// It also returns true if there are positions that contain a space indicating that the game is still in play
