package main

import (
	"fmt"

	"github.com/finecoding/oxo-1/oxo"
)

func main() {

	/*  Setting up a game.

	A Game contains a slice of Turns. A Turn contains the 3 x 3 board in Grid.
	We need to configure Player X and O each with a Tactic function.

	Take note of the data types in type.go.
	A Board is a field in a Turn which is an element of a slice of Turns in a Game.
	Game also has two Player type fields for O and X that each have a Tactic method that
	returns an integer for the position to place the their marker.

	Start by appending a Turn with one empty Grid onto the Game
	We only have one tactic method for a Player at the moment that just chooses and empty space at random.

	This is a standalone program.

	It will be later expanded to become a command line program that will accept configuration.


	*/

	// Game is the highest level data structure.  Players, Boards and Turns are embedded fields.
	var g oxo.Game
	// configuration
	g.O.Tactic = oxo.CompleteXLine
	g.X.Tactic = oxo.Corner
	g.O.Name = "COMPLETEXLINE"
	g.X.Name = "CORNER"
    //We now have tactics for RANDOM, CORNER, CENTRE, COMPLETEXLINE and COMPLETEOLINE
	//Start with a blank grid where every position is a space.
	g.Turns = append(g.Turns, oxo.Turn{Board: oxo.Grid([]byte("         "))})

	// Now we need to find the state of the Game at any turn.  Do this by using the lookup table calculated in status.go
	findstate := oxo.Newlookup()
	// flip between O and X
	var flip bool
	// Game loop.  It is never going to be more than 9 turns.  A minimum Game is 5 Turns
	// Note that X goes first by default.  
	// This might give an advantage and so we should make this configurable later.

	for turn := 0; turn < 9; turn++ {

		if flip {
			pos := g.O.Tactic(g.Turns[turn].Board) //uses the Tactic function to obtain an integer between 0 and 8
			g.Turns[turn].Board[pos] = oxo.O       //O updates the Board for this Turn

		} else {
			pos := g.X.Tactic(g.Turns[turn].Board)
			g.Turns[turn].Board[pos] = oxo.X
		}

		//Convert the Board, which is a 9 byte array into a slice and into a string
		//using the Grid2string method on Grid.
		//Thi uses findstates function that uses the big lookup map to find a string holding the state of the board.
		//This state string is used to update the status field for this Turn in the Game.
		//The status is used immediately to exit the loop if the state of the board is no longer in play following the player
		//updating the board with their move.

		g.Turns[turn].Status = findstate[g.Turns[turn].Board.Grid2string()]
		//If the Status is not in Play, the Game is over and we exit the loop
		if g.Turns[turn].Status != "PLAY" {
			break
		}
		// Otherwise we append another Turn to the Game Turn slice
		g.Turns = append(g.Turns, g.Turns[turn])
		// And we flip from O to X or vice versa
		flip = !flip
		// The loop around
	}
	// Finally print out the game across the terminal using the Game stringer method.
	// Best results if Jetbrains Mono font is used in the terminal
	// If a font is used that does not contain box drawing symbols, it is not going to look good!
	fmt.Println(g)

}
