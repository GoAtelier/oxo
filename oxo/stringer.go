package oxo

import (
	"fmt"
)

// String methods for each of the main data types
//
//String method for Grid making it compatible with Stringer interface used by fmt.
//Returns a formatted string that can be printed be used by fmt to print the 3x3 grid on the terminal

func (b Grid) String() string {

	//New Improved version with box drawing.  Requires a font like Jetbrains Mono that has symbols for drawing boxes

	l1 := fmt.Sprintf("┌───┬───┬───┐\n")
	l2 := fmt.Sprintf("│ %s│ %s│ %s│\n", string(b[0])+" ", string(b[1])+" ", string(b[2])+" ")
	l3 := fmt.Sprintf("├───┼───┼───┤\n")
	l4 := fmt.Sprintf("│ %s│ %s│ %s│\n", string(b[3])+" ", string(b[4])+" ", string(b[5])+" ")
	l5 := fmt.Sprintf("├───┼───┼───┤\n")
	l6 := fmt.Sprintf("│ %s│ %s│ %s│\n", string(b[6])+" ", string(b[7])+" ", string(b[8])+" ")
	l7 := fmt.Sprintf("└───┴───┴───┘\n")

	return fmt.Sprintf("%s%s%s%s%s%s%s\n\n", l1, l2, l3, l4, l5, l6, l7)

}

func (b Grid) Grid2string() string {

	//A utility method to help streamline the game loop

	return string(b[:])

}

//String method for Turn making it compatible with Stringer interface used by fmt.
//Returns a formatted string that can be printed be used by fmt to print the 3x3 grid on the terminal
//With an extra line to hold the status of the Turn (XWIN, OWIN, DRAW, PLAY)

func (t Turn) String() string {
	l1 := fmt.Sprintf("┌───┬───┬───┐\n")
	l2 := fmt.Sprintf("│ %s│ %s│ %s│\n", string(t.Board[0])+" ", string(t.Board[1])+" ", string(t.Board[2])+" ")
	l3 := fmt.Sprintf("├───┼───┼───┤\n")
	l4 := fmt.Sprintf("│ %s│ %s│ %s│\n", string(t.Board[3])+" ", string(t.Board[4])+" ", string(t.Board[5])+" ")
	l5 := fmt.Sprintf("├───┼───┼───┤\n")
	l6 := fmt.Sprintf("│ %s│ %s│ %s│\n", string(t.Board[6])+" ", string(t.Board[7])+" ", string(t.Board[8])+" ")
	l7 := fmt.Sprintf("└───┴───┴───┘\n")
	l8 := fmt.Sprintf(" %s", string(t.Status))

	return fmt.Sprintf("%s%s%s%s%s%s%s%s\n\n", l1, l2, l3, l4, l5, l6, l7, l8)
}

/*
	String method for Game making it compatible with Stringer interface used by fmt.

Game has a slice of Turns, each of which have a Grid and Status string.
Game also has two Player fields for O and X
String method ranges over the Turns slice printing between 5 and 9 3 x 3 Grids across the screen with the Status under each.
String also the names of the two players.  It returns one long string with formatting to print a completed Game on multiple lines.
*/
func (g Game) String() string {
	var l1, l2, l3, l4, l5, l6, l7, l8, l9 string

	for i := range g.Turns {
		//	l1 = l1 + string(g.Turns[i].Board[0]) + string(g.Turns[i].Board[1]) + string(g.Turns[i].Board[2]) + "    "
		//	l2 = l2 + string(g.Turns[i].Board[3]) + string(g.Turns[i].Board[4]) + string(g.Turns[i].Board[5]) + "    "
		//	l3 = l3 + string(g.Turns[i].Board[6]) + string(g.Turns[i].Board[7]) + string(g.Turns[i].Board[8]) + "    "
		//	l4 = l4 + g.Turns[i].Status + "   "
		l1 = l1 + fmt.Sprintf("┌───┬───┬───┐")
		l2 = l2 + fmt.Sprintf("│ %s│ %s│ %s│", string(g.Turns[i].Board[0])+" ", string(g.Turns[i].Board[1])+" ", string(g.Turns[i].Board[2])+" ")
		l3 = l3 + fmt.Sprintf("├───┼───┼───┤")
		l4 = l4 + fmt.Sprintf("│ %s│ %s│ %s│", string(g.Turns[i].Board[3])+" ", string(g.Turns[i].Board[4])+" ", string(g.Turns[i].Board[5])+" ")
		l5 = l5 + fmt.Sprintf("├───┼───┼───┤")
		l6 = l6 + fmt.Sprintf("│ %s│ %s│ %s│", string(g.Turns[i].Board[6])+" ", string(g.Turns[i].Board[7])+" ", string(g.Turns[i].Board[8])+" ")
		l7 = l7 + fmt.Sprintf("└───┴───┴───┘")
		l8 = l8 + g.Turns[i].Status + "         "

	}
	l9 = l9 + "Players: " + g.O.Name + " v " + g.X.Name

	return fmt.Sprintf("\n\n\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n     %s\n\n                 %s\n", l1, l2, l3, l4, l5, l6, l7, l8, l9)

}
