package main

import (
	"fmt"
)

type grid [9]byte

func (board grid) String() string {

	l1 := string(board[0]) + string(board[1]) + string(board[2])
	l2 := string(board[3]) + string(board[4]) + string(board[5])
	l3 := string(board[6]) + string(board[7]) + string(board[8])
	return fmt.Sprintf("\n%s\n%s\n%s\n\n", l1, l2, l3)
}

func main() {

	fmt.Println("hi")

	var board = grid([]byte("   OOOXXX"))

	//Here we have a 9 byte array created from a byte slice created by converting a string literal
	//Note, I could have done this anonymously using the [9]byte([]byte("XXXOOOXXX")) instead of grid([]byte("XXXOOOXXX"))
	//But the named type grid will come in useful later when I want to write some methods around it.

	fmt.Println("here is....", board)
	// fmt.Println("here it is as string again....", string([]byte(board)))
	// Interestingly, I seem to be able to convert a string to a slice of bytes and then to an array of 9 bytes.
	// But I cannot take an array of 9 bytes and convert it back into a string.
	// There must be a trick to doing this.

	// https: //golangdocs.com/golang-byte-array-to-string

	// Now the Stringer method seems to work fine with an array instead of a slice.
	// It matches the method set of the stringer interface and so can be used with fmt.

	// I could make that String method much more elaborate, for example put it all in a little line grid.
	// However. I will save that until we print out a whole game across the terminal!

	fmt.Printf("%s", board)

}
