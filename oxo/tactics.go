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
	if g[Mc] == S {
		return 4
	} else {
		return Random(g)
	}
}

// Note the use of constants for grid positions in the 3 x 3 board.  Tl=Topright=0, Tc=Topcentre=1, Tr=Topright=2....see types.go
// Corner tactic selects any vacant corner otherwise it randomly selects any other empty location
func Corner(g Grid) int {

	switch {
	case g[Tl] == S:
		return 0
		fallthrough
	case g[Tr] == S:
		return 2
		fallthrough
	case g[Bl] == S:
		return 6
		fallthrough
	case g[Br] == S:
		return 8
		fallthrough
	default:
		return Random(g)
	}

}
func CompleteOLine(g Grid) int {
	r := 9
	switch {
	//top row
	case (g[Tl] == O && g[Tc] == O) && g[Tr] == S:
		r = 2
		fallthrough
	case (g[Tc] == O && g[Tr] == O) && g[Tl] == S:
		r = 0
		fallthrough
	case (g[Tl] == O && g[Tr] == O) && g[Tc] == S:
		r = 1
		fallthrough

		//mid row
	case (g[Ml] == O && g[Mc] == O) && g[Mr] == S:
		r = 5
		fallthrough
	case (g[Ml] == O && g[Mr] == O) && g[Mc] == S:
		r = 4
		fallthrough
	case (g[Mc] == O && g[Mr] == O) && g[Ml] == S:
		r = 3
		fallthrough

		//bottom row
	case (g[Bl] == O && g[Bc] == O) && g[Br] == S:
		r = 8
		fallthrough
	case (g[Bc] == O && g[Br] == O) && g[Bl] == S:
		r = 6
		fallthrough
	case (g[Bl] == O && g[Br] == O) && g[Bc] == S:
		r = 7
		fallthrough

		//left col
	case (g[Tl] == O && g[Ml] == O) && g[Bl] == S:
		r = 6
		fallthrough
	case (g[Tl] == O && g[Bl] == O) && g[Ml] == S:
		r = 3
		fallthrough
	case (g[Ml] == O && g[Bl] == O) && g[Tl] == S:
		r = 0
		fallthrough

		//mid col
	case (g[Tc] == O && g[Mc] == O) && g[Bc] == S:
		r = 7
		fallthrough
	case (g[Mc] == O && g[Bc] == O) && g[Tc] == S:
		r = 1
		fallthrough
	case (g[Tc] == O && g[Bc] == O) && g[Mc] == S:
		r = 4
		fallthrough

		// right col
	case (g[Tr] == O && g[Mr] == O) && g[Br] == S:
		r = 8
		fallthrough
	case (g[Tr] == O && g[Br] == O) && g[Mr] == S:
		r = 5
		fallthrough
	case (g[Mr] == O && g[Br] == O) && g[Tr] == S:
		r = 2
		fallthrough

		// top left diag
	case (g[Tl] == O && g[Mc] == O) && g[Bc] == S:
		r = 7
		fallthrough
	case (g[Mc] == O && g[Bc] == O) && g[Tl] == S:
		r = 0
		fallthrough
	case (g[Tl] == O && g[Bc] == O) && g[Mc] == S:
		r = 4
		fallthrough

		// top right diag
	case (g[Tr] == O && g[Mc] == O) && g[Bl] == S:
		r = 6
		fallthrough
	case (g[Tr] == O && g[Bl] == O) && g[Mc] == S:
		r = 4
		fallthrough
	case (g[Mc] == O && g[Bl] == O) && g[Tr] == S:
		r = 2
		fallthrough
	default:
		r = Random(g)
	}
	return r
	// /
}
func CompleteXLine(g Grid) int {
	r := 9
	switch {
	//top row
	case (g[Tl] == X && g[Tc] == X) && g[Tr] == S:
		r = 2
		fallthrough
	case (g[Tc] == X && g[Tr] == X) && g[Tl] == S:
		r = 0
		fallthrough
	case (g[Tl] == X && g[Tr] == X) && g[Tc] == S:
		r = 1
		fallthrough

		//mid row
	case (g[Ml] == X && g[Mc] == X) && g[Mr] == S:
		r = 5
		fallthrough
	case (g[Ml] == X && g[Mr] == X) && g[Mc] == S:
		r = 4
		fallthrough
	case (g[Mc] == X && g[Mr] == X) && g[Ml] == S:
		r = 3
		fallthrough

		//bottom row
	case (g[Bl] == X && g[Bc] == X) && g[Br] == S:
		r = 8
		fallthrough
	case (g[Bc] == X && g[Br] == X) && g[Bl] == S:
		r = 6
		fallthrough
	case (g[Bl] == X && g[Br] == X) && g[Bc] == S:
		r = 7
		fallthrough

		//left col
	case (g[Tl] == X && g[Ml] == X) && g[Bl] == S:
		r = 6
		fallthrough
	case (g[Tl] == X && g[Bl] == X) && g[Ml] == S:
		r = 3
		fallthrough
	case (g[Ml] == X && g[Bl] == X) && g[Tl] == S:
		r = 0
		fallthrough

		//mid col
	case (g[Tc] == X && g[Mc] == X) && g[Bc] == S:
		r = 7
		fallthrough
	case (g[Mc] == X && g[Bc] == X) && g[Tc] == S:
		r = 1
		fallthrough
	case (g[Tc] == X && g[Bc] == X) && g[Mc] == S:
		r = 4
		fallthrough

		// right col
	case (g[Tr] == X && g[Mr] == X) && g[Br] == S:
		r = 8
		fallthrough
	case (g[Tr] == X && g[Br] == X) && g[Mr] == S:
		r = 5
		fallthrough
	case (g[Mr] == X && g[Br] == X) && g[Tr] == S:
		r = 2
		fallthrough

		// top left diag
	case (g[Tl] == X && g[Mc] == X) && g[Bc] == S:
		r = 7
		fallthrough
	case (g[Mc] == X && g[Bc] == X) && g[Tl] == S:
		r = 0
		fallthrough
	case (g[Tl] == X && g[Bc] == X) && g[Mc] == S:
		r = 4
		fallthrough

		// top right diag
	case (g[Tr] == X && g[Mc] == X) && g[Bl] == S:
		r = 6
		fallthrough
	case (g[Tr] == X && g[Bl] == X) && g[Mc] == S:
		r = 4
		fallthrough
	case (g[Mc] == X && g[Bl] == X) && g[Tr] == S:
		r = 2
		fallthrough

	default:
		r = Random(g)
	}
	return r
}
