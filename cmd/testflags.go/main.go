package main

import (
	"fmt"
	// "github.com/finecoding/oxo-1/oxo"
)
type GameStateflags uint8
//type AllStateflags uint8

//GameStateflags is what we need to test for the Game loop

const (
	isInplay GameStateflags = 1 << iota
	isOwin
	isXwin
	isLegal
	//naming the AllStateflags - capitalisation indicates the state of the flag l = legal, x = xwin, o = owin, p=in play
	lxop0000  GameStateflags =  iota //ILLEGAL      
	lxoP0001  //ILLEGAL      
	lxOp0010  //ILLEGAL      
	lxOP0011  //ILLEGAL     
	lXop0100  //ILLEGAL      
	lXoP0101  //ILLEGAL      
	lXOp0110  //ILLEGAL      
	lXOP0111  //ILLEGAL
	Lxop1000  //DRAW  Legal board Neither O or X win and game is full
	LxoP1001  //PLAY  Legal board Neither O or X win and game is in play
	LxOp1010  //OWIN  Legal board OWins and game is full
	LxOP1011  //OWIN  Legal board OWins and game is in play
	LXop1100  //XWIN  Legal board XWins and game is full
	LXoP1101  //XWIN  Legal board XWins and game is in play
	LXOp1110  //ERROR Legal board but has both X AND O winning lines and is full      
	LXOP1111  //ERROR Legal board but has both X and O winning line and is in play    



)



func Set(b, flag GameStateflags) GameStateflags    { return b | flag }
	func Clear(b, flag GameStateflags) GameStateflags  { return b &^ flag }
	func Toggle(b, flag GameStateflags) GameStateflags { return b ^ flag }
	func Has(b, flag GameStateflags) bool    { return b&flag != 0 }
	

func main() {

    //AllStateflags is a set of 16 mutually exclusive states for a 3x3 board containing O,X or Space at each position
	//They will contain many illegal states.  The legal states overlap and some should never be reached 
	//For example a full board which would be a draw cannot have a winning line of X's or O's.
	//AllStateflags captures where several states are true at the same time.  So we can decide what must go into our GameState lookup table.
    //



	





/* 	So these 16 flags are the combinations of the 4 states we need for the game.
	First we need a map of all game states string[flags]
	This first map will be populated by concentric loops to generate a 9 char string for
    the key, which will be parsed by functions to identify which flags should be set.
	We will create a second map of string[Stateflags] to use as a lookup table for the state
	of the game.  We will set each value with a case switch statement that examines each key in the string[flags] map 
    rewrite this!!
 */

	fmt.Printf("lxop0000\t\t\t%08b\n", lxop0000)
	fmt.Printf("lxoP0001\t\t\t%08b\n", lxoP0001)
	fmt.Printf("lxOp0010\t\t\t%08b\n", lxOp0010)
	fmt.Printf("lxOP0011\t\t\t%08b\n", lxOP0011)
	fmt.Printf("lXop0100\t\t\t%08b\n", lXop0100)
	fmt.Printf("lXoP0101\t\t\t%08b\n", lXoP0101)
	fmt.Printf("lXOp0110\t\t\t%08b\n", lXOp0110)
	fmt.Printf("lXOP0111\t\t\t%08b\n", lXOP0111)
	fmt.Printf("Lxop1000\t\t\t%08b\n", Lxop1000)
	fmt.Printf("LxoP1001\t\t\t%08b\n", LxoP1001)
	fmt.Printf("LxOp1010\t\t\t%08b\n", LxOp1010)
	fmt.Printf("LxOP1011\t\t\t%08b\n", LxOP1011)
	fmt.Printf("LXop1100\t\t\t%08b\n", LXop1100)
	fmt.Printf("LXoP1101\t\t\t%08b\n", LXoP1101)
	fmt.Printf("LXOp1110\t\t\t%08b\n", LXOp1110)
	fmt.Printf("LXOP1111\t\t\t%08b\n", LXOP1111)




	fmt.Println() 



	fmt.Printf("isInplay\t\t\t%08b\n", isInplay)
	fmt.Printf("isOwin\t\t\t\t%08b\n", isXwin)
	fmt.Printf("isXwin\t\t\t\t%08b\n", isOwin)
	fmt.Printf("isLegal\t\t\t\t%08b\n", isLegal)

	fmt.Println("Hi")

	fmt.Printf("-------------------------\n")
	var b GameStateflags

	b = Set(b, isInplay)
	fmt.Printf("After set isInplay\t\t%08b\n", isInplay)


	b = Set(b, isOwin)
	fmt.Printf("After set isOwin\t\t%08b\n", isOwin)
	b = Set(b, isXwin)
	fmt.Printf("After set isXwin\t\t%08b\n", isXwin)

	b = Set(b, isLegal)
	fmt.Printf("After set isLegal\t\t%08b\n", isLegal)

	
	fmt.Printf("Anding them together....\t%08b\n", isInplay|isLegal|isXwin|isOwin)
}
