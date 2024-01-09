// Data types go here as the need arises to define them.

package oxo

import (
	"math/rand"
)
type GameStateclass uint8 //Holds 16 permutations of 4 GameStates 

type GameState uint8 //Holds 4 GameStates OWIN,XWIN,INPLAY and LEGAL

//These are the 4 GameStates, a Board can have multiple GameStates 
const (
	IsInplay GameState = 1 << iota
	IsOwin
	IsXwin
	IsLegal
)
//These are all 16 permutations of 4 GameStates so there are NO overlaps.  Any board is in only one of these.
const (
	Classlxop0000 GameStateclass = iota //ILLEGAL
	ClasslxoP0001                       //ILLEGAL
	ClasslxOp0010                       //ILLEGAL
	ClasslxOP0011                       //ILLEGAL
	Classlxop0100                       //ILLEGAL
	ClasslxoP0101                       //ILLEGAL
	ClasslxOp0110                       //ILLEGAL
	ClasslxOP0111                       //ILLEGAL
	Classlxop1000                       //DRAW  Legal board Neither O or X win and game is full
	ClasslxoP1001                       //PLAY  Legal board Neither O or X win and game is in play
	ClasslxOp1010                       //OWIN  Legal board OWins and game is full.  OWin, not a DRAW
	ClasslxOP1011                       //OWIN  Legal board OWins and game is in play
	Classlxop1100                       //XWIN  Legal board XWins and game is full.  XWin, not a DRAW
	ClasslxoP1101                       //XWIN  Legal board XWins and game is in play
	ClasslxOp1110                       //ERROR Legal board but has both X AND O winning lines and is full. Error because one should be detected first.
	ClasslxOP1111                       //ERROR Legal board but has both X and O winning line and is in play.  Error because one should be detected first.
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
