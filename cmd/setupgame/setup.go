package main

import (
	"fmt"

	"github.com/finecoding/oxo-1/oxo"
)

func main() {

	/*  Setting up a game.

	type Game struct {
		Turns  []Turn
		X      Player
		O      Player
		Result string

		type Turn struct {
		Board  Grid
		Status string
	}
	type Player struct {
		Name   string
		Tactic func(Grid) int
		Rank   int

	A Game contains a slice of Turns. A Turn contains the 3 x 3 board in Grid.
	We need to configure Player X and O each with a Tactic function.

	Start by appending a Turn with one empty Grid onto the Game
	The Player names are a long form version.

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

		//Convert the Board, which is a 9 byte slice into a string
		stb := [9]byte(g.Turns[turn].Board)
		str := string([]byte(stb[:]))
		//find the state of the game using the lookup table and update the Status field of the Turn
		g.Turns[turn].Status = findstate[str]
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
