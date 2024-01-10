package oxo

import (
	"math/rand"
)

// Random is a tactic returning an int that is its chosen location on the grid
// Findspaces finds the spaces on the Grid and Random choses one....randomly
func Random(g Grid) int {
	spc, _ := g.FindSpaces()
	return int(spc[randInt(0, len(spc))])
}

// randInt choses a random number between two values
func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Centre tactic selects the centre square if it is a space. Otherwise it randomly selects any other empty location
func Centre(g Grid) int {
	if g[4] == S {
		return 4
	} else {
		return Random(g)
	}
}

// Corner tactic selects any vacant corner otherwise it randomly selects any other empty location
func Corner(g Grid) int {

	switch {
	case g[0] == S:
		return 0
		fallthrough
	case g[2] == S:
		return 2
		fallthrough
	case g[6] == S:
		return 6
		fallthrough
	case g[8] == S:
		return 8
		fallthrough
	default:
		return Random(g)
	}

}
func CompleteOLine(g Grid) int {
    r:=9
	switch {
	//top row
	case (g[0] == O && g[1] == O) && g[2] == S:
		r = 2
		fallthrough
	case (g[1] == O && g[2] == O) && g[0] == S:
		r = 0
		fallthrough
	case (g[0] == O && g[2] == O) && g[1] == S:
		r = 1
		fallthrough

		//mid row
	case (g[3] == O && g[4] == O) && g[5] == S:
		r = 5
		fallthrough
	case (g[3] == O && g[5] == O) && g[4] == S:
		r = 4
		fallthrough
	case (g[4] == O && g[5] == O) && g[3] == S:
		r = 3
		fallthrough

		//bottom row
	case (g[6] == O && g[7] == O) && g[8] == S:
		r = 8
		fallthrough
	case (g[7] == O && g[8] == O) && g[6] == S:
		r = 6
		fallthrough
	case (g[6] == O && g[8] == O) && g[7] == S:
		r = 7
		fallthrough

		//left col
	case (g[0] == O && g[3] == O) && g[6] == S:
		r = 6
		fallthrough
	case (g[0] == O && g[6] == O) && g[3] == S:
		r = 3
		fallthrough
	case (g[3] == O && g[6] == O) && g[0] == S:
		r = 0
		fallthrough

		//mid col
	case (g[1] == O && g[4] == O) && g[7] == S:
		r = 7
		fallthrough
	case (g[4] == O && g[7] == O) && g[1] == S:
		r = 1
		fallthrough
	case (g[1] == O && g[7] == O) && g[4] == S:
		r = 4
		fallthrough

		// right col
	case (g[2] == O && g[5] == O) && g[8] == S:
		r = 8
		fallthrough
	case (g[2] == O && g[8] == O) && g[5] == S:
		r = 5
		fallthrough
	case (g[5] == O && g[8] == O) && g[2] == S:
		r = 2
		fallthrough

		// top left diag
	case (g[0] == O && g[4] == O) && g[7] == S:
		r = 7
		fallthrough
	case (g[4] == O && g[7] == O) && g[0] == S:
		r = 0
		fallthrough
	case (g[0] == O && g[7] == O) && g[4] == S:
		r = 4
		fallthrough

		// top right diag
	case (g[2] == O && g[4] == O) && g[6] == S:
		r = 6
		fallthrough
	case (g[2] == O && g[6] == O) && g[4] == S:
		r = 4
		fallthrough
	case (g[4] == O && g[6] == O) && g[2] == S:
		r = 2

	default:
		r = Random(g)
	}
return r
	// /
}
func CompleteXLine(g Grid) int {
	r:=9
	switch {
	//top row
	case (g[0] == X && g[1] == X) && g[2] == S:
		r = 2
		fallthrough
	case (g[1] == X && g[2] == X) && g[0] == S:
		r = 0
		fallthrough
	case (g[0] == X && g[2] == X) && g[1] == S:
		r = 1
		fallthrough

		//mid row
	case (g[3] == X && g[4] == X) && g[5] == S:
		r = 5
		fallthrough
	case (g[3] == X && g[5] == X) && g[4] == S:
		r = 4
		fallthrough
	case (g[4] == X && g[5] == X) && g[3] == S:
		r = 3
		fallthrough

		//bottom row
	case (g[6] == X && g[7] == X) && g[8] == S:
		r = 8
		fallthrough
	case (g[7] == X && g[8] == X) && g[6] == S:
		r = 6
		fallthrough
	case (g[6] == X && g[8] == X) && g[7] == S:
		r = 7
		fallthrough

		//left col
	case (g[0] == X && g[3] == X) && g[6] == S:
		r = 6
		fallthrough
	case (g[0] == X && g[6] == X) && g[3] == S:
		r = 3
		fallthrough
	case (g[3] == X && g[6] == X) && g[0] == S:
		r = 0
		fallthrough

		//mid col
	case (g[1] == X && g[4] == X) && g[7] == S:
		r = 7
		fallthrough
	case (g[4] == X && g[7] == X) && g[1] == S:
		r = 1
		fallthrough
	case (g[1] == X && g[7] == X) && g[4] == S:
		r = 4
		fallthrough

		// right col
	case (g[2] == X && g[5] == X) && g[8] == S:
		r = 8
		fallthrough
	case (g[2] == X && g[8] == X) && g[5] == S:
		r = 5
		fallthrough
	case (g[5] == X && g[8] == X) && g[2] == S:
		r = 2
		fallthrough

		// top left diag
	case (g[0] == X && g[4] == X) && g[7] == S:
		r = 7
		fallthrough
	case (g[4] == X && g[7] == X) && g[0] == S:
		r = 0
		fallthrough
	case (g[0] == X && g[7] == X) && g[4] == S:
		r = 4
		fallthrough

		// top right diag
	case (g[2] == X && g[4] == X) && g[6] == S:
		r = 6
		fallthrough
	case (g[2] == X && g[6] == X) && g[4] == S:
		r = 4
		fallthrough
	case (g[4] == X && g[6] == X) && g[2] == S:
		r = 2
		fallthrough

	default:
		r = Random(g)
	}
	return r
}
