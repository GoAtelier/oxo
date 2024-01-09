package oxo

// Newlookup returns a lookup table that has a 9 character string as key that represents the OXO board and value string that indicates its state
func Newlookup() map[string]string {

	//This is our big lookup table for all combinations of a 3 x 3 grid where each position contains either an X, O or space
	//That is 3 to the power 9 = 19685 permutations
	var def, ILLEGALlxop0000, ILLEGALlxoP0001, ILLEGALlxOp0010, ILLEGALlxOP0011, ILLEGALlXop0100, ILLEGALlXoP0101, ILLEGALlXOp0110, ILLEGALlXOP0111, DRAWLxop1000, PLAYLxoP1001, OWINLxOp1010, OWINLxOP1011, XWINLXop1100, XWINLXoP1101, ERRORLXOp1110, ERRORLXOP1111 int
	//vars used as counters for testing.
	gsl := make(map[string]string)

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
											// from numbers 0,1 and 2 to letters O,X and SPACE using the convert function
											s := convert(p0) + convert(p1) + convert(p2) + convert(p3) + convert(p4) + convert(p5) + convert(p6) + convert(p7) + convert(p8)
											// s holds a nine character string of characters O,X and SPACE
											// gs is a byte of type Gamestateclass.  The set functions set positional bits according
											// to the state of the board.
											gs := SetXwin(s) | SetOwin(s) | SetLegal(s) | SetPlay(s)
											// This is for testing the clauses of the following switch case statement
											// A count helps to verify the expected number of each category of board state.
											//											if !Has(gs,isInplay) && !Has(gs, isXwin)&& Has(gs, isOwin)&& Has(gs,isLegal) {
											//												c++
											//											fmt.Printf("gs is [%08b] s = [%s] \n",gs,s)
											//}
											// The positional bits need to be shifted 4 bits to the to right to be used in a switch
											// not needed gs = gs >> 4
											// So we have sixteen different combinations of four states that a board can be in.
											// Base on whether X or O has a winning line, the board is full or not and whether it is
											// in a legal state ie. the max difference between the number of X's and O's is one
											// The case switch identifies combinations that are legal, but also impossible in a game.
											// For example having a both O and X win at the same time.  One would always be detected first.
											// It also identifies board states where a Draw (because all places are full) might also have
											// and O or X winning line.
											//
											// The switch case statement allows for 16 mutually exclusive combinations and allows the board
											// state to be labelled with the appropriate string.
											// The result is a single lookup table stored in map gsl the has maps a string representation of
											// board to a string the indicates its state.
											//
											//The result is game state lookup table (gsl) which will be used later in the game loop.
											//Each clause has a counter for testing

											switch gs {
											case 0b00000000:
												gsl[s] = "ILLEGAL"//-lxop0000"
												ILLEGALlxop0000++
												break
											case 0b00000001:
												gsl[s] = "ILLEGAL"//-lxoP0001" //ILLEGAL
												ILLEGALlxoP0001++
												break
											case 0b00000010:
												gsl[s] = "ILLEGAL"//-lxOp0010" //ILLEGAL
												ILLEGALlxOp0010++
												break
											case 0b00000011:
												gsl[s] = "ILLEGAL"//-lxOP0011" //ILLEGAL
												ILLEGALlxOP0011++
												break
											case 0b00000100:
												gsl[s] = "ILLEGAL"//-lXop0100" //ILLEGAL
												ILLEGALlXop0100++
												break
											case 0b00000101:
												gsl[s] = "ILLEGAL"//-lXoP0101" //ILLEGAL
												ILLEGALlXoP0101++
												break
											case 0b00000110:
												gsl[s] = "ILLEGAL"//-lXOp0110" //ILLEGAL
												ILLEGALlXOp0110++
											case 0b00000111:
												gsl[s] = "ILLEGAL"//-lXOP0111" //ILLEGAL
												ILLEGALlXOP0111++
												break
											case 0b00001000:
												gsl[s] = "DRAW"//-Lxop1000" //DRAW  Legal board Neither O or X win and game is full
												DRAWLxop1000++
												break
											case 0b00001001:
												gsl[s] = "PLAY"//-LxoP1001" PLAY  Legal board Neither O or X win and game is in play
												PLAYLxoP1001++
												break
											case 0b00001010:
												gsl[s] = "OWIN"//-LxOp1010" OWIN  Legal board OWins and game is full
												OWINLxOp1010++
												break
											case 0b00001011:
												gsl[s] = "OWIN"//-LxOP1011" OWIN  Legal board OWins and game is in play
												OWINLxOP1011++
												break
											case 0b00001100:
												gsl[s] = "XWIN"//-LXop1100" //XWIN  Legal board XWins and game is full
												XWINLXop1100++
												break
											case 0b00001101:
												gsl[s] = "XWIN"//-LXoP1101" //XWIN  Legal board XWins and game is in play
												XWINLXoP1101++
												break
											case 0b00001110:
												gsl[s] = "ERROR"//-LXOp1110" //ERROR Legal board but has both X AND O winning lines and is full
												//In normal play one win will be detected for either O or X, there will be both!
												ERRORLXOp1110++
												break
											case 0b00001111:
												gsl[s] = "ERROR"//-LXOP1111" //ERROR Legal board has both X and O winning but is in play
												ERRORLXOP1111++
												break
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
	return gsl //a map of ["O O OXXXX"]"XWIN"
}

// Used to build a string for the key to the lookup map
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
