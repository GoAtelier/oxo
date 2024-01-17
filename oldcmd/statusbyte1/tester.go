package main

import (
	"fmt"
	"github.com/finecoding/oxo-1/oxo"
)
func main () {

	fmt.Println("Hi")
	r:=oxo.SetXwin("O O   XXX")
	fmt.Printf("SetXwin\t\t\t\t%08b\n", r)
	r=oxo.SetXwin("XXX   OO ")
	fmt.Printf("SetXwin\t\t\t\t%08b\n", r)
	r=oxo.SetXwin("OO XXXO  ")
	fmt.Printf("SetXwin\t\t\t\t%08b\n", r)
	r=oxo.SetXwin("X  X  XOO")
	fmt.Printf("SetXwin\t\t\t\t%08b\n", r)
	r=oxo.SetXwin(" X  X OXO")
	fmt.Printf("SetXwin\t\t\t\t%08b\n", r)
	r=oxo.SetXwin("OOX  X  X")
	fmt.Printf("SetXwin\t\t\t\t%08b\n", r)
	r=oxo.SetXwin("X   XO  X")
	fmt.Printf("SetXwin\t\t\t\t%08b\n", r)
	r=oxo.SetXwin("  X X XOO")
	fmt.Printf("SetXwin\t\t\t\t%08b\n", r)
	r=oxo.SetXwin("         ")
	fmt.Printf("SetXwin\t\t\t\t%08b\n", r)
	r=oxo.SetXwin("OOOOOOOOO")
	fmt.Printf("SetXwin\t\t\t\t%08b\n", r)

	r=oxo.SetXwin("XXXXXXXXX")
	r=oxo.SetXwin("         ")
	fmt.Printf("SetXwin\t\t\t\t%08b\n", r)
	fmt.Printf("IsXWin is \t\t\t%08b\n", oxo.IsXwin)
    fmt.Println(r==oxo.IsXwin)
	



//	fmt.Println(r)
}