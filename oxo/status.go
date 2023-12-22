// This is a standalone command used in development.
// It uses concentric loops to iterate over every combination of states of a 3 x 3 grid containing either O, X or a space at each location
// It creates one very large lookup table thats has a string of 9 characters representing the 3 x 3 grid as its key and a status string as its value
// This will be tidied up an added to the oxo package, then used in the game loop to lookup the state of a game when the board is updated
// after each turn taken by a player.
// It detects O win, X win, a draw or an illegal board state.
//

package oxo

import (
	
	"strings"
)

func convert(j int) string {
	switch j {
	case 0:
		return "O"
	case 1:
		return "X"
	case 2:
		return " "
	}
	return "?"
}

// Transpose takes the byte slice version of the string representing the 3 x 3 grid and transposes it
// to make it easier usable by a string.Contains function
func transpose(sb []byte) string {

	return string(sb[0]) + string(sb[3]) + string(sb[6]) + string(sb[1]) + string(sb[4]) + string(sb[7]) + string(sb[2]) + string(sb[5]) + string(sb[8])

}

// Markstatus takes a nine character string and returns a string that indicates the status of the game
// The alternatives are "XWIN", "OWIN", "DRAW" or "ILLEGAL"

func legality(s string) string {
	x := 88
	o := 79
	O := byte(o)
	X := byte(x)
	sb := []byte(s)
	r := "LEGAL"

	// Illegal states
	// If there is both row of three X's and a line three O's.
	// One or the other would win before this state could exist, so is it illegal

	switch {
	case sb[0] == O && sb[1] == O && sb[2] == O && sb[3] == X && sb[4] == X && sb[5] == X:
		r = "ERROR"
		fallthrough
	case sb[0] == X && sb[1] == X && sb[2] == X && sb[3] == O && sb[4] == O && sb[5] == O:
		r = "ERROR"
		fallthrough
	case sb[3] == O && sb[4] == O && sb[5] == O && sb[6] == X && sb[7] == X && sb[8] == X:
		r = "ERROR"
		fallthrough
	case sb[3] == X && sb[4] == X && sb[5] == X && sb[6] == O && sb[7] == O && sb[8] == O:
		r = "ERROR"

	}

	// More illegal game states.  There can only be a difference of 1 between the number of X's and O's

	if (strings.Count(s, "X") - strings.Count(s, "O")) > 1 {
		r = "ERROR"
	}
	if (strings.Count(s, "O") - strings.Count(s, "X")) > 1 {

		r = "ERROR"
	}

	return r
}

func addstate(legal map[string]string) map[string]string{
	x := 88
	o := 79
	O := byte(o)
	X := byte(x)
	r := "LEGAL"
	//   s:="         "
	//	sb := []byte(s)
	 trx, mrx, brx, fcx, scx, tcx, ldx, rdx := 0, 0, 0, 0, 0, 0, 0, 0
	 tro, mro, bro, fco, sco, tco, ldo, rdo := 0, 0, 0, 0, 0, 0, 0, 0
	length := 0
	for s := range legal {
		length++
		sb := []byte(s)

		if strings.Contains(s, " ") {
			r = "PLAY"
			
		} else {
			r = "DRAW"
		
		}

		// Note a draw means each of the 9 locations have either a O or an X.  So that is 2 power 9 = 512 permutations

		switch {

		case sb[0] == X && sb[1] == X && sb[2] == X: //top row
			trx++
			r = "XWIN"
			fallthrough
		case sb[3] == X && sb[4] == X && sb[5] == X: //middle row
			r = "XWIN"
			mrx++
			fallthrough
		case sb[6] == X && sb[7] == X && sb[8] == X: //bottom row
			r = "XWIN"
			brx++
			fallthrough
		case sb[0] == X && sb[3] == X && sb[6] == X: //first col
			r = "XWIN"
			fcx++
			fallthrough
		case sb[1] == X && sb[4] == X && sb[7] == X: //second col
			r = "XWIN"
			scx++
			fallthrough
		case sb[2] == X && sb[5] == X && sb[8] == X: //third col
			r = "XWIN"
			tcx++
			fallthrough
		case sb[0] == X && sb[4] == X && sb[8] == X: //left diag
			r = "XWIN"
			ldx++
			fallthrough
		case sb[2] == X && sb[4] == X && sb[6] == X: //right diag
			r = "XWIN"
			rdx++

		}
		switch {
		case sb[0] == O && sb[1] == O && sb[2] == O:
			r = "OWIN"
			tro++
			fallthrough
		case sb[3] == O && sb[4] == O && sb[5] == O:
			r = "OWIN"
			mro++
			fallthrough
		case sb[6] == O && sb[7] == O && sb[8] == O:
			r = "OWIN"
			bro++
			fallthrough
		case sb[0] == O && sb[3] == O && sb[6] == O:
			r = "OWIN"
			fco++
			fallthrough
		case sb[1] == O && sb[4] == O && sb[7] == O:
			r = "OWIN"
			sco++
			fallthrough
		case sb[2] == O && sb[5] == O && sb[8] == O:
			r = "OWIN"
			tco++
			fallthrough
		case sb[0] == O && sb[4] == O && sb[8] == O:
			r = "OWIN"
			ldo++
			fallthrough
		case sb[2] == O && sb[4] == O && sb[6] == O:
			r = "OWIN"
			rdo++
		}
		legal[s] = r

	}
//	fmt.Printf(" length of legal %d\n spins of legal %d\n", len(legal), length)
//	fmt.Printf("trx=%d mrx=%d brx=%d fcx=%d scx=%d tcx=%d ldx=%d rdx=%d total=%d\n", trx, mrx, brx, fcx, scx, tcx, ldx, rdx, trx+mrx+brx+fcx+scx+tcx+ldx+rdx)
//	fmt.Printf("tro=%d mro=%d bro=%d fco=%d sco=%d tco=%d ldo=%d rdo=%d total=%d\n", tro, mro, bro, fco, sco, tco, ldo, rdo, tro+mro+bro+fco+sco+tco+ldo+rdo)
	return legal

}
func Newlookup() (map[string]string, map[string]string) {

	//This is our big lookup table for all combinations of a 3 x 3 grid where each position contains either an X, O or space
	//That is 3 to the power 9 = 19685 permutations

	legallookup := make(map[string]string)
	illegallookup := make(map[string]string)

	//p is the position in a nine character string
	//to create 19685 strings we iterate using numbers 0 to 2 to represent the three values.
	//then use the convert function to return a string 0="O", 1="X" and 2=" "
	for p0 := 0; p0 < 3; p0++ {
		for p1 := 0; p1 < 3; p1++ {
			for p2 := 0; p2 < 3; p2++ {
				for p3 := 0; p3 < 3; p3++ {
					for p4 := 0; p4 < 3; p4++ {
						for p5 := 0; p5 < 3; p5++ {
							for p6 := 0; p6 < 3; p6++ {
								for p7 := 0; p7 < 3; p7++ {

									for p8 := 0; p8 < 3; p8++ {
										{
											s := convert(p0) + convert(p1) + convert(p2) + convert(p3) + convert(p4) + convert(p5) + convert(p6) + convert(p7) + convert(p8)
											// s holds a nine character string which becomes the key
											// function markstatus derives the state of the game we assign it to the value

											if legality(s) == "LEGAL" {
												legallookup[s] = legality(s)
											} else {
												illegallookup[s] = legality(s)
											}
										}
									}
								}

							}
						}
					}
				}
			}
		}
	}

	return addstate(legallookup), illegallookup
}
