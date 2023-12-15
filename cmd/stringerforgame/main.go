package main

import (
	"fmt"

	"github.com/finecoding/oxo-1/oxo"
)

// This command introduces a String method to make the Game type printable using fmt.
//
// Grid represents the state of a turn of the game on a 3 x 3 grid
// Turn adds a status string
// Game is a slice of Turns
//
// A Game contains enough data to print a complete game of between 5 and 9 turns across the screen

// This makes the Grid type conform to the Stringer interface so it can be used with fmt.
// It takes the 3 x 3 stored as an array in Grid returns a string.

// This makes the Game type conform to the Stringer interface so it can be used with fmt.
// It loops over the Game slice, building three strings from the contents of each board.
// It combines these to return a string, which can be used by fmt.

func main() {

	// Here we are building a set of literally defined turns.

	var t1 = oxo.Turn{
		Board:  oxo.Grid([]byte("X........")),
		Status: "PLAY",
	}
	var t2 = oxo.Turn{
		Board:  oxo.Grid([]byte("X.O......")),
		Status: "PLAY",
	}
	var t3 = oxo.Turn{
		Board:  oxo.Grid([]byte("X.OX.....")),
		Status: "PLAY",
	}
	var t4 = oxo.Turn{
		Board:  oxo.Grid([]byte("X.OXO....")),
		Status: "PLAY",
	}
	var t5 = oxo.Turn{
		Board:  oxo.Grid([]byte("X.OXO.X..")),
		Status: "XWIN",
	}
	var g oxo.Game

	// Here we append all the turns into a game slice.

	g = append(g, t1, t2, t3, t4, t5)

	// Thanks to the String method, it is easy to print out the game
	fmt.Println(g)
	fmt.Println(t1.Board)
	// So far we have used literals to build a game slice.  Good for testing the String method
	// The next development will be to get the game loop working to build the game slice
	// go run oxo2.go

	// This is what it prints:

	//  X..    X.O    X.O    X.O    X.O
	//  ...    ...    X..    XO.    XO.
	//  ...    ...    ...    ...    X..
	//
	//  PLAY   PLAY   PLAY   PLAY   XWIN
}
