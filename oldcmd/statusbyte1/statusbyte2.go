package main

import (
	"fmt"
	"github.com/finecoding/oxo-1/oxo"
)

func main() {

	fmt.Printf("Classlxop0000\t\t\t%08b\n", oxo.Classlxop0000)
	fmt.Printf("ClasslxoP0001\t\t\t%08b\n", oxo.ClasslxoP0001)
	fmt.Printf("ClasslxOp0010\t\t\t%08b\n", oxo.ClasslxOp0010)
	fmt.Printf("ClasslxOP0011\t\t\t%08b\n", oxo.ClasslxOP0011)
	fmt.Printf("Classlxop0100\t\t\t%08b\n", oxo.Classlxop0100)
	fmt.Printf("ClasslxoP0101\t\t\t%08b\n", oxo.ClasslxoP0101)
	fmt.Printf("ClasslxOp0110\t\t\t%08b\n", oxo.ClasslxOp0110)
	fmt.Printf("ClasslxOP0111\t\t\t%08b\n\n", oxo.ClasslxOP0111)
	fmt.Printf("Classlxop1000\t\t\t%08b\n", oxo.Classlxop1000)
	fmt.Printf("ClasslxoP1001\t\t\t%08b\n", oxo.ClasslxoP1001)
	fmt.Printf("ClasslxOp1010\t\t\t%08b\n", oxo.ClasslxOp1010)
	fmt.Printf("ClasslxOP1011\t\t\t%08b\n", oxo.ClasslxOP1011)
	fmt.Printf("Classlxop1100\t\t\t%08b\n", oxo.Classlxop1100)
	fmt.Printf("ClasslxoP1101\t\t\t%08b\n", oxo.ClasslxoP1101)
	fmt.Printf("ClasslxOp1110\t\t\t%08b\n", oxo.ClasslxOp1110)
	fmt.Printf("ClasslxOP1111\t\t\t%08b\n", oxo.ClasslxOP1111)

	fmt.Println()

	fmt.Printf("IsLegal\t\t\t\t%08b\n", oxo.IsLegal)
	fmt.Printf("IsOwin\t\t\t\t%08b\n", oxo.IsXwin)
	fmt.Printf("IsXwin\t\t\t\t%08b\n", oxo.IsOwin)
	fmt.Printf("IsInplay\t\t\t%08b\n", oxo.IsInplay)

	fmt.Println("Hi")

	fmt.Printf("-------------------------\n")

	gsl := oxo.Newlookup()

	fmt.Println()
	fmt.Printf("[OOOXXXOOO] is %s\n", gsl["OOOXXXOOO"])
	fmt.Printf("[      OOO] is %s\n", gsl["      OOO"])
	fmt.Printf("[    XXOOO] is %s\n", gsl["    XXOOO"])
	fmt.Printf("[XXX   OOO] is %s\n", gsl["XXX   OOO"])
	fmt.Printf("[XX     OO] is %s\n", gsl["XX     OO"])
	fmt.Printf("[X      OO] is %s\n", gsl["X      OO"])
	fmt.Printf("[XXX  OOOO] is %s\n", gsl["XXX  OOOO"])

}
