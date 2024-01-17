/* Total 19683
XWIN 896
OWIN 896
DRAW 132
ILLEGAL 17759
SUM 19683 */

package main

import (

	"fmt"
	"github.com/finecoding/oxo-1/oxo"
)

func main() {

	
	goodstate,badstate :=oxo.Newlookup()
	

	
//	fmt.Printf("Total %d\nXWIN %d \nOWIN %d \nDRAW %d \nILLEGAL %d\nLEGAL %d\nPLAY %d \nSUM %d\n", len(checkstate), xwins, owins, draws, illegal,legal, play, xwins+owins+illegal+draws+play+legal)
fmt.Printf("\n\nLEGAL %d\nILLEGAL %d \nSUM %d\n", len(goodstate), len(badstate), len(goodstate)+len(badstate))

xwins, owins, draws, play := 0, 0, 0, 0
for i, value := range goodstate {
	switch value {

	case "XWIN":
		xwins++
	case "OWIN":
		owins++
	case "DRAW":
		draws++	
		fmt.Println(oxo.Grid([]byte(i)))
	case "PLAY":
		play++
	}

}
fmt.Printf("Total %d\nXWIN %d \nOWIN %d \nDRAW %d \nPLAY %d\nSUM %d\n", len(goodstate), xwins, owins, draws,play,xwins+owins+draws+play)




}