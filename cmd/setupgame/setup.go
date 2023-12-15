package main

import (
	"fmt"

	"github.com/finecoding/oxopfork/oxo"
)

func main() {

	var g oxo.Game
	g.Result = "anything"
	g.O.Tactic = oxo.Random
	g.X.Tactic = oxo.Random
	g.Turns = []oxo.Turn{}
	g.Turns = append(g.Turns, oxo.Turn{Board: oxo.Grid([]byte("          "))})
	g.Turns[0].Status = "PLAY"
	g.O.Name = "RANDOM"
	g.O.Tactic = oxo.Random
	g.X.Tactic = oxo.Random
	g.X.Name = "RANDOM"
	// We have the Game setup.  Now make a loop
//	pos := g.O.Tactic(g.Turns[0].Board)
//	g.Turns[0].Board[pos] = byte(79) //O = 78 X=88 SPC=21
//	fmt.Println(pos)
//	fmt.Println(g.Turns[0].Board[0],g.Turns[0].Board[2],g.Turns[0].Board[3],g.Turns[0].Board[4],g.Turns[0].Board[5],g.Turns[0].Board[6],g.Turns[0].Board[7],g.Turns[0].Board[8])
//	fmt.Println(g.Turns[0].Board)
var flip bool
	for turn:=0; turn < 9; turn++ {
		// fmt.Println(turn)
		if flip {
		pos := g.O.Tactic(g.Turns[turn].Board)
		g.Turns[turn].Board[pos] = byte(79)
		} else {
		pos := g.X.Tactic(g.Turns[turn].Board)
		g.Turns[turn].Board[pos] = byte(88)
		}
		g.Turns = append(g.Turns, g.Turns[turn])
		flip=!flip
			}
	fmt.Println(g)
		}
