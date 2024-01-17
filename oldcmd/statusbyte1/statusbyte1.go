package main

import (
	"fmt"
	"strings"
	// "github.com/finecoding/oxo-1/oxo"
	// "github.com/finecoding/oxo-1/oxo"
)

type GameStateflags uint8

//type GameStateflags uint8

//GameStateflags is what we need to test for the Game loop

/*
	 const (
		isInplay GameStateflags = 1 << iota
		isOwin
		isXwin
		isLegal
		//naming the GameStateflags - capitalisation indicates the state of the flag l = legal, x = xwin, o = owin, p=in play
		lxop0000 GameStateflags = iota //ILLEGAL
		lxoP0001                       //ILLEGAL
		lxOp0010                       //ILLEGAL
		lxOP0011                       //ILLEGAL
		lXop0100                       //ILLEGAL
		lXoP0101                       //ILLEGAL
		lXOp0110                       //ILLEGAL
		lXOP0111                       //ILLEGAL
		Lxop1000                       //DRAW  Legal board Neither O or X win and game is full
		LxoP1001                       //PLAY  Legal board Neither O or X win and game is in play
		LxOp1010                       //OWIN  Legal board OWins and game is full
		LxOP1011                       //OWIN  Legal board OWins and game is in play
		LXop1100                       //XWIN  Legal board XWins and game is full
		LXoP1101                       //XWIN  Legal board XWins and game is in play
		LXOp1110                       //ERROR Legal board but has both X AND O winning lines and is full
		LXOP1111                       //ERROR Legal board but has both X and O winning line and is in play
*/
const (
	isInplay = 0b10000000
	isOwin   = 0b01000000
	isXwin   = 0b00100000
	isLegal  = 0b00010000
	// naming the GameStateflags - capitalisation indicates the state of the flag l = legal, x = xwin, o = owin, p=in play
	lxop0000 = 0b00000000 //ILLEGAL
	lxoP0001 = 0b00000001 //ILLEGAL
	lxOp0010 = 0b00000010 //ILLEGAL
	lxOP0011 = 0b00000011 //ILLEGAL
	lXop0100 = 0b00000100 //ILLEGAL
	lXoP0101 = 0b00000101 //ILLEGAL
	lXOp0110 = 0b00000110 //ILLEGAL
	lXOP0111 = 0b00000111 //ILLEGAL
	Lxop1000 = 0b00001000 //DRAW  Legal board Neither O or X win and game is full
	LxoP1001 = 0b00001001 //PLAY  Legal board Neither O or X win and game is in play
	LxOp1010 = 0b00001010 //OWIN  Legal board OWins and game is full
	LxOP1011 = 0b00001011 //OWIN  Legal board OWins and game is in play
	LXop1100 = 0b00001100 //XWIN  Legal board XWins and game is full
	LXoP1101 = 0b00001101 //XWIN  Legal board XWins and game is in play
	LXOp1110 = 0b00001110 //ERROR Legal board but has both X AND O winning lines and is full
	LXOP1111 = 0b00001111 //ERROR Legal board but has both X and O winning line and is in play
)

func Set(b, flag GameStateflags) GameStateflags    { return b | flag }
func Clear(b, flag GameStateflags) GameStateflags  { return b &^ flag }
func Toggle(b, flag GameStateflags) GameStateflags { return b ^ flag }
func Has(b, flag GameStateflags) bool              { return b&flag != 0 }
 
func setXwin(s string) GameStateflags {
	var r GameStateflags
	r=0b00000000
	x := 88
	X := byte(x)
	sb := []byte(s)
	switch {

	case (sb[0] == X && sb[1] == X && sb[2] == X): //top row
	//	fmt.Println("found XXX top row")
		r = isXwin
		fallthrough
	case sb[3] == X && sb[4] == X && sb[5] == X: //middle row
		r = isXwin
		fallthrough
	case sb[6] == X && sb[7] == X && sb[8] == X: //bottom row
		r = isXwin
		fallthrough
	case sb[0] == X && sb[3] == X && sb[6] == X: //first col
		r = isXwin
		fallthrough
	case sb[1] == X && sb[4] == X && sb[7] == X: //second col
		r = isXwin
		fallthrough
	case sb[2] == X && sb[5] == X && sb[8] == X: //third col
		r = isXwin
		fallthrough
	case sb[0] == X && sb[4] == X && sb[8] == X: //left diag
		r = isXwin
		fallthrough
	case sb[2] == X && sb[4] == X && sb[6] == X: //right diag
		r = isXwin

	}
	return r
}
func setPlay(sb string) GameStateflags {
	var r GameStateflags
	if strings.Count(sb, " ") > 0 {
		r = isInplay 

	}
	return r
}

func setOwin(s string) GameStateflags {
	var r GameStateflags
	o := 79
	O := byte(o)
	sb := []byte(s)
	switch {
	case sb[0] == O && sb[1] == O && sb[2] == O:
		r = isOwin
		fallthrough
	case sb[3] == O && sb[4] == O && sb[5] == O:
		r = isOwin
		fallthrough
	case sb[6] == O && sb[7] == O && sb[8] == O:
		r = isOwin
		fallthrough
	case sb[0] == O && sb[3] == O && sb[6] == O:
		r = isOwin
		fallthrough
	case sb[1] == O && sb[4] == O && sb[7] == O:
		r = isOwin
		fallthrough
	case sb[2] == O && sb[5] == O && sb[8] == O:
		r = isOwin
		fallthrough
	case sb[0] == O && sb[4] == O && sb[8] == O:
		r = isOwin
		fallthrough
	case sb[2] == O && sb[4] == O && sb[6] == O:
		r = isOwin
	}

	return r
}

func setLegal(s string) GameStateflags {

	var r GameStateflags
	// More illegal game states.  There can only be a difference of 1 between the number of X's and O's
	diff:=strings.Count(s, "X") - strings.Count(s, "O")
	if diff<2 && diff>-2 {
		r = isLegal
	}

	return r
}
	
	
	



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

/* func setlookup(gs GameStateflags, gsl map[string]string, s string) {
var int	ILLEGALlxop0000,ILLEGALlxoP0001,ILLEGALlxOp0010,ILLEGALlxOP0011,ILLEGALlXop0100,ILLEGALlXoP0101,ILLEGALlXOp0110,ILLEGALlXOP0111,DRAWLxop1000,PLAYLxoP1000,OWINLxOp1010,OWINLxOP1010,XWINLXop1100,XWINLXoP1101,ERRORLXOp1110,ERRORLXOP1111

	switch gs {
	case lxop0000:
		gsl[s] = "ILLEGAL-lxop0000"
		ILLEGALlxop0000++
	case lxoP0001:
		gsl[s] = "ILLEGAL-lxoP0001" //ILLEGAL
		ILLEGAL-lxoP0001++
	case lxOp0010:
		gsl[s] = "ILLEGAL-lxOp0010" //ILLEGAL
		ILLEGAL-lxOp0010++
	case lxOP0011:
		gsl[s] = "ILLEGAL-lxOP0011" //ILLEGAL
		ILLEGAL-lxOP0011++
	case lXop0100:
		gsl[s] = "ILLEGAL-lXop0100" //ILLEGAL
		ILLEGAL-lXop0100++
	case lXoP0101:
		gsl[s] = "ILLEGAL-lXoP0101" //ILLEGAL
		ILLEGAL-lXoP0101++
	case lXOp0110:
		gsl[s] = "ILLEGAL-lXOp0110" //ILLEGAL
		ILLEGAL-lXOp0110++
	case lXOP0111:
		gsl[s] = "ILLEGAL-lXOP0111" //ILLEGAL
		ILLEGAL-lXOP0111++
	case Lxop1000:
		gsl[s] = "DRAW-Lxop1000" //DRAW  Legal board Neither O or X win and game is full
		DRAW-Lxop1000++
	case LxoP1001:
		gsl[s] = "PLAY-LxoP1000" //PLAY  Legal board Neither O or X win and game is in play
		PLAY-LxoP1000++
	case LxOp1010:
		gsl[s] = "OWIN-LxOp1010" //OWIN  Legal board OWins and game is full
		OWIN-LxOp1010++
	case LxOP1011:
		gsl[s] = "OWIN-LxOP1010" //OWIN  Legal board OWins and game is in play
		OWIN-LxOP1010++
	case LXop1100:
		gsl[s] = "XWIN-LXop1100" //XWIN  Legal board XWins and game is full
		XWIN-LXop1100++
	case LXoP1101:
		gsl[s] = "XWIN-LXoP1101" //XWIN  Legal board XWins and game is in play
		XWIN-LXoP1101++
	case LXOp1110:
		gsl[s] = "ERROR-LXOp1110" //ERROR Legal board but has both X AND O winning lines and is full
		ERROR-LXOp1110++
	case LXOP1111:
		gsl[s] = "ERROR-LXOP1111" //ERROR L
		ERROR-LXOP1111++
	}
	return
} */

func Newlookup() map[string]string {

	//This is our big lookup table for all combinations of a 3 x 3 grid where each position contains either an X, O or space
	//That is 3 to the power 9 = 19685 permutations
	var def, ILLEGALlxop0000, ILLEGALlxoP0001, ILLEGALlxOp0010, ILLEGALlxOP0011, ILLEGALlXop0100, ILLEGALlXoP0101, ILLEGALlXOp0110, ILLEGALlXOP0111, DRAWLxop1000, PLAYLxoP1001, OWINLxOp1010, OWINLxOP1011, XWINLXop1100, XWINLXoP1101, ERRORLXOp1110, ERRORLXOP1111 int
	//var reg GameStateflags
	gsl := make(map[string]string)
	c:=0
	//	AllStateLookup := make(map[string]string)
	//var r GameStateflags

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
											// s holds a nine character string which becomes the s
											// function markstatus derives the state of the game we assign it to the value
											//	gs := (setPlay(s) |setOwin(s)|setXwin(s) |setLegal(s))
											//fmt.Printf("s= [%s] and gs=[%08b] \n",s,gs)
											//	gs := setPlay(s)
											//	gs := setOwin(s,r)
											gs := setXwin(s) | setOwin(s) | setLegal(s) | setPlay(s)
											
											//fmt.Printf("s = %s Xwin and Owin %08b\n",s,gs)
											
											//	gs := setLegal(s)
											
											if !Has(gs,isInplay) && !Has(gs, isXwin)&& Has(gs, isOwin)&& Has(gs,isLegal) {
												c++
											fmt.Printf("gs is [%08b] s = [%s] \n",gs,s)
											}
											
											
											gs=gs>>4
											switch gs {
											case 0b00000000:
												gsl[s] = "ILLEGAL-lxop0000"
												ILLEGALlxop0000++
												break
											//	fallthrough
											case 0b00000001:
												gsl[s] = "ILLEGAL-lxoP0001" //ILLEGAL
												ILLEGALlxoP0001++
												break
											//	fallthrough
											case 0b00000010:
												gsl[s] = "ILLEGAL-lxOp0010" //ILLEGAL
												ILLEGALlxOp0010++
												break
											//	fallthrough
											case 0b00000011:
												gsl[s] = "ILLEGAL-lxOP0011" //ILLEGAL
												ILLEGALlxOP0011++
												break
											//	fallthrough
											case 0b00000100:
												gsl[s] = "ILLEGAL-lXop0100" //ILLEGAL
												ILLEGALlXop0100++
												break
											//	fallthrough
											case 0b00000101:
												gsl[s] = "ILLEGAL-lXoP0101" //ILLEGAL
												ILLEGALlXoP0101++
												break
											case 0b00000110:
												gsl[s] = "ILLEGAL-lXOp0110" //ILLEGAL
												ILLEGALlXOp0110++
											//	fallthrough
											case 0b00000111:
												gsl[s] = "ILLEGAL-lXOP0111" //ILLEGAL
												ILLEGALlXOP0111++
												break
											//	fallthrough
											case 0b00001000:
												gsl[s] = "DRAW-Lxop1000" //DRAW  Legal board Neither O or X win and game is full
												DRAWLxop1000++
												break
											//	fallthrough
											case 0b00001001:
												gsl[s] = "PLAY-LxoP1001" //PLAY  Legal board Neither O or X win and game is in play
												PLAYLxoP1001++
												break
											//	fallthrough
											case 0b00001010:
												gsl[s] = "OWIN-LxOp1010" //OWIN  Legal board OWins and game is full
												OWINLxOp1010++
												break
											case 0b00001011:
												gsl[s] = "OWIN-LxOP1011" //OWIN  Legal board OWins and game is in play
												OWINLxOP1011++
												break
											//	fallthrough
											case 0b00001100:
												gsl[s] = "XWIN-LXop1100" //XWIN  Legal board XWins and game is full
												XWINLXop1100++
												break
												
											//	fallthrough
											case 0b00001101:
												gsl[s] = "XWIN-LXoP1101" //XWIN  Legal board XWins and game is in play
												XWINLXoP1101++
												break
											//	fmt.Printf("gs is %0b s = [%s] INPLAY\n",gs,s)
											//	fallthrough
											case 0b00001110:
												gsl[s] = "ERROR-LXOp1110" //ERROR Legal board but has both X AND O winning lines and is full
												//In normal play one win will be detected for either O or X, there will be both! 	
												ERRORLXOp1110++
												break
											//	fallthrough
											case 0b00001111:
												gsl[s] = "ERROR-LXOP1111" //ERROR Legal board has both X and O winning but is in play
												ERRORLXOP1111++
												break
											//	fallthrough
											default:
												def++

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
    fmt.Printf("%d counted \n",c)
	fmt.Printf("ILLEGALlxop0000=%d\n", ILLEGALlxop0000)
	fmt.Printf("ILLEGALlxoP0001=%d\n", ILLEGALlxoP0001)
	fmt.Printf("ILLEGALlxOp0010=%d\n", ILLEGALlxOp0010)
	fmt.Printf("ILLEGALlxOP0011=%d\n", ILLEGALlxOP0011)
	fmt.Printf("ILLEGALlXop0100=%d\n", ILLEGALlXop0100)
	fmt.Printf("ILLEGALlXoP0101=%d\n", ILLEGALlXoP0101)
	fmt.Printf("ILLEGALlXOp0110=%d\n", ILLEGALlXOp0110)
	fmt.Printf("ILLEGALlXOP0111=%d\n", ILLEGALlXOP0111)
	fmt.Printf("DRAWLxop1000=%d\n", DRAWLxop1000)
	fmt.Printf("PLAYLxoP1001=%d\n", PLAYLxoP1001)
	fmt.Printf("OWINLxOp1010=%d\n", OWINLxOp1010)
	fmt.Printf("OWINLxOP1011=%d\n", OWINLxOP1011)
	fmt.Printf("XWINLXop1100=%d\n", XWINLXop1100)
	fmt.Printf("XWINLXoP1101=%d\n", XWINLXoP1101)
	fmt.Printf("ERRORLXOp1110=%d\n", ERRORLXOp1110)
	fmt.Printf("ERRORLXOP1111=%d\n", ERRORLXOP1111)
	fmt.Printf("Total %d\n", ILLEGALlxop0000+ILLEGALlxoP0001+ILLEGALlxOp0010+ILLEGALlxOP0011+ILLEGALlXop0100+ILLEGALlXoP0101+ILLEGALlXOp0110+ILLEGALlXOP0111+DRAWLxop1000+PLAYLxoP1001+OWINLxOp1010+OWINLxOP1011+XWINLXop1100+XWINLXoP1101+ERRORLXOp1110+ERRORLXOP1111)
	fmt.Println("Should total 19,683!")
	fmt.Printf("Default = %d ", def)
	fmt.Printf("Default+Others = %d ", def+ILLEGALlxop0000+ILLEGALlxoP0001+ILLEGALlxOp0010+ILLEGALlxOP0011+ILLEGALlXop0100+ILLEGALlXoP0101+ILLEGALlXOp0110+ILLEGALlXOP0111+DRAWLxop1000+PLAYLxoP1001+OWINLxOp1010+OWINLxOP1011+XWINLXop1100+XWINLXoP1101+ERRORLXOp1110+ERRORLXOP1111)

	return gsl //a map of ["O O OXXXX"]"XWIN"
}

func main() {

	fmt.Printf("lxop0000\t\t\t%08b\n", lxop0000)
	fmt.Printf("lxoP0001\t\t\t%08b\n", lxoP0001)
	fmt.Printf("lxOp0010\t\t\t%08b\n", lxOp0010)
	fmt.Printf("lxOP0011\t\t\t%08b\n", lxOP0011)
	fmt.Printf("lXop0100\t\t\t%08b\n", lXop0100)
	fmt.Printf("lXoP0101\t\t\t%08b\n", lXoP0101)
	fmt.Printf("lXOp0110\t\t\t%08b\n", lXOp0110)
	fmt.Printf("lXOP0111\t\t\t%08b\n\n", lXOP0111)
	fmt.Printf("Lxop1000\t\t\t%08b\n", Lxop1000)
	fmt.Printf("LxoP1001\t\t\t%08b\n", LxoP1001)
	fmt.Printf("LxOp1010\t\t\t%08b\n", LxOp1010)
	fmt.Printf("LxOP1011\t\t\t%08b\n", LxOP1011)
	fmt.Printf("LXop1100\t\t\t%08b\n", LXop1100)
	fmt.Printf("LXoP1101\t\t\t%08b\n", LXoP1101)
	fmt.Printf("LXOp1110\t\t\t%08b\n", LXOp1110)
	fmt.Printf("LXOP1111\t\t\t%08b\n", LXOP1111)

	fmt.Println()

	fmt.Printf("isLegal\t\t\t\t%08b\n", isLegal)
	fmt.Printf("isOwin\t\t\t\t%08b\n", isXwin)
	fmt.Printf("isXwin\t\t\t\t%08b\n", isOwin)
	fmt.Printf("isInplay\t\t\t%08b\n", isInplay)

	fmt.Println("Hi")

	fmt.Printf("-------------------------\n")

	gsl := Newlookup()
//	for key,value :=range gsl {
//if value 

//	}

	fmt.Println()
	fmt.Printf("[OOOXXXOOO] is %s\n",gsl["OOOXXXOOO"])
	fmt.Printf("[      OOO] is %s\n",gsl["      OOO"]) 
	fmt.Printf("[    XXOOO] is %s\n",gsl["    XXOOO"])
	fmt.Printf("[XXX   OOO] is %s\n",gsl["XXX   OOO"])
	fmt.Printf("[XX     OO] is %s\n",gsl["XX     OO"])
	fmt.Printf("[X      OO] is %s\n",gsl["X      OO"])
	fmt.Printf("[XXX  OOOO] is %s\n",gsl["XXX  OOOO"])
	// Now we need to test the gsl works!
	/*    c:=0
		for s, state := range gsl {
	    if state=="DRAW-Lxop1000" {
			c++
			fmt.Printf("%s\t\t\t%s\n", oxo.Grid([]byte(s)), state)
		}
	}*/
	//fmt.Printf("count DRAWLxop100 = %d",c)

}
