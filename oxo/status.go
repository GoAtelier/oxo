package oxo

import (
	"strings"
)

//These functions examine a string representation of the 3 x 3 board and derive its Gamestate
//The Gamestate is single byte and the state is one of four bit positions defined by constants IsXwin,IsOwin, IsInplay and IsLegal
//These will be used to compose Gamestateclass. 

func SetXwin(s string) GameState {
	var r GameState
	X := byte(88)
	//Convert to bytes to access the index
	sb := []byte(s)
	switch {

	case (sb[0] == X && sb[1] == X && sb[2] == X): //top row
		r = IsXwin
		fallthrough
	case sb[3] == X && sb[4] == X && sb[5] == X: //middle row
		r = IsXwin
		fallthrough
	case sb[6] == X && sb[7] == X && sb[8] == X: //bottom row
		r = IsXwin
		fallthrough
	case sb[0] == X && sb[3] == X && sb[6] == X: //first col
		r = IsXwin
		fallthrough
	case sb[1] == X && sb[4] == X && sb[7] == X: //second col
		r = IsXwin
		fallthrough
	case sb[2] == X && sb[5] == X && sb[8] == X: //third col
		r = IsXwin
		fallthrough
	case sb[0] == X && sb[4] == X && sb[8] == X: //left diag
		r = IsXwin
		fallthrough
	case sb[2] == X && sb[4] == X && sb[6] == X: //right diag
		r = IsXwin

	}
	return r
}
func SetPlay(sb string) GameState {
	var r GameState
	if strings.Count(sb, " ") > 0 {
		r = IsInplay

	}
	return r
}

func SetOwin(s string) GameState {
	var r GameState
	O := byte(79)
	//Convert to bytes to access the index
	sb := []byte(s)
	switch {
	case sb[0] == O && sb[1] == O && sb[2] == O:
		r = IsOwin
		fallthrough
	case sb[3] == O && sb[4] == O && sb[5] == O:
		r = IsOwin
		fallthrough
	case sb[6] == O && sb[7] == O && sb[8] == O:
		r = IsOwin
		fallthrough
	case sb[0] == O && sb[3] == O && sb[6] == O:
		r = IsOwin
		fallthrough
	case sb[1] == O && sb[4] == O && sb[7] == O:
		r = IsOwin
		fallthrough
	case sb[2] == O && sb[5] == O && sb[8] == O:
		r = IsOwin
		fallthrough
	case sb[0] == O && sb[4] == O && sb[8] == O:
		r = IsOwin
		fallthrough
	case sb[2] == O && sb[4] == O && sb[6] == O:
		r = IsOwin
	}

	return r
}

func SetLegal(s string) GameState {

	var r GameState
	// Illegal game states.  There can only be a difference of 1 between the number of X's and O's
	diff := strings.Count(s, "X") - strings.Count(s, "O")
	if diff < 2 && diff > -2 {
		r = IsLegal
	}

	return r
}