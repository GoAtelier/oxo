package oxo

import (
	"strings"
)

//Setxxxx functions examine a string representation of the 3 x 3 board and derive its Gamestate
//The Gamestate is single byte and the state is one of four bit positions defined by constants IsXwin,IsOwin, IsInplay and IsLegal
//These will be used to compose Gamestateclass.

func SetXwin(s string) GameState {
	var r GameState

	sb := []byte(s)
	switch {

	case (sb[Tl] == X && sb[Tc] == X && sb[Tr] == X): //top row
		r = IsXwin
		fallthrough
	case sb[Ml] == X && sb[Mc] == X && sb[Mr] == X: //middle row
		r = IsXwin
		fallthrough
	case sb[Bl] == X && sb[Bc] == X && sb[Br] == X: //bottom row
		r = IsXwin
		fallthrough
	case sb[Tl] == X && sb[Ml] == X && sb[Bl] == X: //left col
		r = IsXwin
		fallthrough
	case sb[Tc] == X && sb[Mc] == X && sb[Bc] == X: //middle col
		r = IsXwin
		fallthrough
	case sb[Tr] == X && sb[Mr] == X && sb[Br] == X: //right col
		r = IsXwin
		fallthrough
	case sb[Tl] == X && sb[Mc] == X && sb[Br] == X: //left diag
		r = IsXwin
		fallthrough
	case sb[Tr] == X && sb[Mc] == X && sb[Bl] == X: //right diag
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
	//Convert to bytes to access the index
	sb := []byte(s)
	switch {

		case (sb[Tl] == O && sb[Tc] == O && sb[Tr] == O): //top row
			r = IsOwin
			fallthrough
		case sb[Ml] == O && sb[Mc] == O && sb[Mr] == O: //middle row
			r = IsOwin
			fallthrough
		case sb[Bl] == O && sb[Bc] == O && sb[Br] == O: //bottom row
			r = IsOwin
			fallthrough
		case sb[Tl] == O && sb[Ml] == O && sb[Bl] == O: //left col
			r = IsOwin
			fallthrough
		case sb[Tc] == O && sb[Mc] == O && sb[Bc] == O: //middle col
			r = IsOwin
			fallthrough
		case sb[Tr] == O && sb[Mr] == O && sb[Br] == O: //right col
			r = IsOwin
			fallthrough
		case sb[Tl] == O && sb[Mc] == O && sb[Br] == O: //left diag
			r = IsOwin
			fallthrough
		case sb[Tr] == O && sb[Mc] == O && sb[Bl] == O: //right diag
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
