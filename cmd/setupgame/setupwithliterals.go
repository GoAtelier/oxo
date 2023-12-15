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
	g.Turns = append(g.Turns, oxo.Turn{Board: oxo.Grid([]byte("O OXXX  X"))})
	g.Turns = append(g.Turns, oxo.Turn{Board: oxo.Grid([]byte("O OXX   X"))})
	g.Turns = append(g.Turns, oxo.Turn{Board: oxo.Grid([]byte("O OX   X "))})
	g.Turns = append(g.Turns, oxo.Turn{Board: oxo.Grid([]byte("O O   X  "))})
	g.Turns = append(g.Turns, oxo.Turn{Board: oxo.Grid([]byte("O OXXX   "))})
	g.Turns[0].Status = "PLAY"
	g.Turns[1].Status = "OWIN"
	g.Turns[2].Status = "XWIN"
	g.Turns[3].Status = "DRAW"
	g.Turns[4].Status = "PLAY"

	g.O.Name = "RANDOM"
	g.X.Name = "RANDOM"

	fmt.Println(g)
}
