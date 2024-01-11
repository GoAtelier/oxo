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

	//We now have tactics for RANDOM, CORNER, CENTRE, COMPLETEXLINE and COMPLETEOLINE
	//Start with a blank grid where every position is a space.

	// Now we need to find the state of the Game at any turn.  Do this by using the lookup table calculated in status.go
	nl := oxo.Newlookup()

	//Now to add these to a new data structure, a Result, which will be a slice of Games.

	reps := 1000
	var res oxo.Group

	for i := 0; i < reps; i++ {

		OPlayer := oxo.Player{"RANDOM", oxo.Corner, 0} //Setup O
		XPlayer := oxo.Player{"RANDOM", oxo.Random, 0} //Setup X
		x := oxo.Playgame(nl, OPlayer, XPlayer, oxo.TossCoin())        //PlayGame, toss coin to decide who goes first...
		res.Games = append(res.Games, x)               //Add Game to res slice.
		res.UpdateNums(x)
		fmt.Println(res.Games[i])
	}
/* 
	fmt.Println(res.Games[0])
	fmt.Println(res.Games[1])
	fmt.Println(res.Games[2]) */

	fmt.Printf("Number of O Wins %d\n", res.NumOwins)
	fmt.Printf("Number of X Wins %d\n", res.NumXwins)
	fmt.Printf("Number of Draws %d\n", res.NumDraws)
	// Next we need to add some file handling to save results in res...

	/* gh repo sync
	   gh repo sync finecoding/oxo-1 --source GoAtelier/oxo
	   gh pr merge -d
	   gh pr create

	   git push --set-upstream origin addresults
	   git commit -m "added constants and results data type"
	   git status -s
	   git add .
	   git checkout addresults
	   git branch addresults
	   git remote -v
	   z oxo-1

	   This, of course, is the reverse history, so start from the bottom going up!
	*/

}
