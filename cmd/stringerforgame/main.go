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
// It loops over the Game slice, building strings from the contents of each board for each line of output.
// It combines these to return a string, which can be used by fmt.

func main() {

	// Here we are building a set of literally defined turns.

	var t1 = oxo.Turn{
		Board:  oxo.Grid([]byte("X        ")),
		Status: "PLAY",
	}
	var t2 = oxo.Turn{
		Board:  oxo.Grid([]byte("X O      ")),
		Status: "PLAY",
	}
	var t3 = oxo.Turn{
		Board:  oxo.Grid([]byte("X OX     ")),
		Status: "PLAY",
	}
	var t4 = oxo.Turn{
		Board:  oxo.Grid([]byte("X OXO    ")),
		Status: "PLAY",
	}
	var t5 = oxo.Turn{
		Board:  oxo.Grid([]byte("X OXO X  ")),
		Status: "XWIN",
	}
	var t6 = oxo.Turn{
		Board:  oxo.Grid([]byte("X OXO XO ")),
		Status: "XWIN",
	}
	var t7 = oxo.Turn{
		Board:  oxo.Grid([]byte("X OXO XOX")),
		Status: "XWIN",
	}
	var t8 = oxo.Turn{
		Board:  oxo.Grid([]byte("X OXOOXOX")),
		Status: "XWIN",
	}
	var t9 = oxo.Turn{
		Board:  oxo.Grid([]byte("XXOXOOXOX")),
		Status: "XWIN",
	}
	var g oxo.Game
	g.X.Name = "RANDOM1"
	g.O.Name = "RANDOM2"
	// Here we append all the turns into a game slice.

	g.Turns = append(g.Turns, t1, t2, t3, t4, t5, t6, t7, t8, t9)

	// Thanks to the String method, it is easy to print out the game
	fmt.Println("Printing a Game")
	fmt.Println(g)
	fmt.Println()
	// This requires a terminal font that has symbols the can be used to create boxes, otherwise the output will be distorted.
	// Jetbrains Mono is one such font.
}

//	┌───┬───┬───┐┌───┬───┬───┐┌───┬───┬───┐┌───┬───┬───┐┌───┬───┬───┐┌───┬───┬───┐┌───┬───┬───┐┌───┬───┬───┐┌───┬───┬───┐
//	│ X │   │   ││ X │   │ O ││ X │   │ O ││ X │   │ O ││ X │   │ O ││ X │   │ O ││ X │   │ O ││ X │   │ O ││ X │ X │ O │
//	├───┼───┼───┤├───┼───┼───┤├───┼───┼───┤├───┼───┼───┤├───┼───┼───┤├───┼───┼───┤├───┼───┼───┤├───┼───┼───┤├───┼───┼───┤
//	│   │   │   ││   │   │   ││ X │   │   ││ X │ O │   ││ X │ O │   ││ X │ O │   ││ X │ O │   ││ X │ O │ O ││ X │ O │ O │
//	├───┼───┼───┤├───┼───┼───┤├───┼───┼───┤├───┼───┼───┤├───┼───┼───┤├───┼───┼───┤├───┼───┼───┤├───┼───┼───┤├───┼───┼───┤
//	│   │   │   ││   │   │   ││   │   │   ││   │   │   ││ X │   │   ││ X │ O │   ││ X │ O │ X ││ X │ O │ X ││ X │ O │ X │
//	└───┴───┴───┘└───┴───┴───┘└───┴───┴───┘└───┴───┴───┘└───┴───┴───┘└───┴───┴───┘└───┴───┴───┘└───┴───┴───┘└───┴───┴───┘
//		 PLAY         PLAY         PLAY         PLAY         XWIN         XWIN         XWIN         XWIN         XWIN
//
//					 Players: RANDOM2 v RANDOM1
