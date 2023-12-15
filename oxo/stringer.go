package oxo

import (
	"fmt"
)

// String methods for each of the main data types
//
//String method for Grid making it compatible with Stringer interface used by fmt.
//Returns a formatted string that can be printed be used by fmt to print the 3x3 grid on the terminal

func (b Grid) String() string {

	l1 := string(b[0]) + string(b[1]) + string(b[2])
	l2 := string(b[3]) + string(b[4]) + string(b[5])
	l3 := string(b[6]) + string(b[7]) + string(b[8])
	return fmt.Sprintf("\n%s\n%s\n%s\n\n", l1, l2, l3)
}

//String method for Turn making it compatible with Stringer interface used by fmt.
//Returns a formatted string that can be printed be used by fmt to print the 3x3 grid on the terminal
//With an extra line to hold the status of the Turn (XWIN, OWIN, DRAW, PLAY)

func (t Turn) String() string {

	l1 := string(t.Board[0]) + string(t.Board[1]) + string(t.Board[2])
	l2 := string(t.Board[3]) + string(t.Board[4]) + string(t.Board[5])
	l3 := string(t.Board[6]) + string(t.Board[7]) + string(t.Board[8])
	l4 := t.Status
	return fmt.Sprintf("\n%s\n%s\n%s\n%s\n\n", l1, l2, l3, l4)
}

// String method for Game making it compatible with Stringer interface used by fmt.
// Returns a formatted string that can be printed be used by fmt to print a row of 3x3 grids on the terminal
// With an extra line to hold the status of the Turn (XWIN, OWIN, DRAW, PLAY)
func (g Game) String() string {
	var l1, l2, l3, l4, l5 string

	for i, _ := range g.Turns {
		l1 = l1 + string(g.Turns[i].Board[0]) + string(g.Turns[i].Board[1]) + string(g.Turns[i].Board[2]) + "    "
		l2 = l2 + string(g.Turns[i].Board[3]) + string(g.Turns[i].Board[4]) + string(g.Turns[i].Board[5]) + "    "
		l3 = l3 + string(g.Turns[i].Board[6]) + string(g.Turns[i].Board[7]) + string(g.Turns[i].Board[8]) + "    "
		l4 = l4 + g.Turns[i].Status + "   "

	}
	l5 = l5 + "Players: "+g.O.Name + " v " + g.X.Name
	return fmt.Sprintf("\n%s\n%s\n%s\n\n%s\n\n%s\n", l1, l2, l3, l4, l5)
}
