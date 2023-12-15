// Data types go here as the need arises to define them.

package oxo

import (
	"math/rand"
)

// board represents the state of a turn of the game on a 3 x 3 grid
type Grid [9]byte

// a turn has a board and status string which will be one of these strings XWIN,OWIN,DRAW,PLAY
type Turn struct {
	Board  Grid
	Status string
}

// a complete game has a minimum of 5 turns and a maximum of 9 turns
// should I set a capacity for the slice?
type Game struct {
	Turns  []Turn
	X      Player
	O      Player
	Result string
}

func (board Grid) FindSpaces() ([]byte, bool) {
	var spc []byte
	for k, v := range board {
		if v == 32 {
			spc = append(spc, byte(k))
		}
	}
	return spc, len(spc) != 0
}

type Player struct {
	Name   string
	Tactic func(Grid) int
	Rank   int
}

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
