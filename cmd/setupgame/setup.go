package main

import (
	"fmt"

	"github.com/finecoding/oxo-1/oxo"
)

func main() {

	/*  Setting up a game.

	A Game contains a slice of Turns. A Turn contains the 3 x 3 board in Grid.
	We need to configure Player X and O each with a Tactic function.

	Start by appending a Turn with one empty Grid onto the Game
	We only have one Player at the moment that just chooses and empty space at random.

	*/

	var g oxo.Game

	g.O.Tactic = oxo.Random
	g.X.Tactic = oxo.Random
	g.Turns = append(g.Turns, oxo.Turn{Board: oxo.Grid([]byte("         "))})
	g.O.Name = "RANDOM"
	g.X.Name = "RANDOM"

	// Now we need to find the state of the Game at any turn.  Do this by using the lookup table calculated in status.go
	findstate := oxo.Newlookup()
	// flip between O and X
	var flip bool
	// Game loop.  It is never going to be more than 9 turns.  A minimum Game is 5 Turns
	for turn := 0; turn < 9; turn++ {

		if flip {
			pos := g.O.Tactic(g.Turns[turn].Board)
			g.Turns[turn].Board[pos] = byte(79) //O

		} else {
			pos := g.X.Tactic(g.Turns[turn].Board)
			g.Turns[turn].Board[pos] = byte(88) //X
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
	fmt.Println(g)

}
