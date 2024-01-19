package oxo

// Newlookup returns a lookup table that has a 9 character string as key that represents the OXO board and value string that indicates its state
func Newlookup() map[string]string {

	//This is our big lookup table for all combinations of a 3 x 3 grid where each position contains either an X, O or space
	//That is 3 to the power 9 = 19685 permutations
	//	var def int
	// used as counter for testing.
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
											// gs is a byte of type Gamestateclass.  These set functions set positional bits according
											// to the state of the board.
											gs := SetXwin(s) | SetOwin(s) | SetLegal(s) | SetPlay(s)
											// So we have sixteen different combinations of four states that a board can be in.
											// Based on whether X or O has a winning line, the board is full or not and whether it is
											// in a legal state ie. the max difference between the number of X's and O's is one
											//
											// This Gamestate byte will be used later to identifies all combinations.
											// Some are legal, but also impossible in a game.
											// For example having a both O and X win with lines of three, at the same time.
											// This is impossible during normal play because one would always be detected first.
											// It also identifies board states where a Draw (because all places are full) might also have
											// and O or X winning line.  These are the edge cases and exceptions that stymie any simple solution.
											// We have to consider all possible combinations of four basic states to decide a state that is
											// consistent with the rules of the game.
											//
											// The gamesstate byte (gs) allows for 16 mutually exclusive combinations of four basic states and this allows the board
											// state to be labelled with the appropriate string.

											// The result is a single lookup table stored in map gsl the has maps a string representation of
											// board to a string the indicates its state.
											//
											//ClassifyGamestate returns a string representing the state of the board, which is used to update the
											//gsl lookup
											gsl[s] = ClassifyGamestate(gs)
											//The result is a complete game state lookup table (gsl).  For any given board ,represented as a key string,
											//it returns the value string for the state of the game.
											//This lookup is central to determining the course of a game and will be used later in the main game loop.
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
	return gsl //gsl, game state lookup, is a map of string to string["O O OXXXX"]"XWIN"
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

// This will be used for creating a Player from a command line string flag
func NewPlayerlookup() map[string]Player {
	pl := make(map[string]Player)
	pl["RANDOM"] = Player{Name: "RANDOM", Tactic: Random, Rank: 0}
	pl["CORNER"] = Player{Name: "CORNER", Tactic: Corner, Rank: 0}
	pl["CENTRE"] = Player{Name: "CENTRE", Tactic: Centre, Rank: 0}
	pl["COMPLETEXLINE"] = Player{Name: "COMPLETEXLINE", Tactic: CompleteXLine, Rank: 0}
	pl["COMPLETEOLINE"] = Player{Name: "COMPLETEOLINE", Tactic: CompleteOLine, Rank: 0}
	return pl
}

//ClassifyGamestate returns a state string that will be used to update the game state lookup map (gls).
//It uses the previously calculated gamestate byte to classify the board into one of 16 mutually exclusive combinations of state.
//It then returns an appropriate state string.

func ClassifyGamestate(gs GameState) string {

	switch GameStateclass(gs) {
	case Classlxop0000:
		return "ILLEGAL" //-lxop0000"
		break
	case ClasslxoP0001:
		return "ILLEGAL" //-lxoP0001" //ILLEGAL
		break
	case ClasslxOp0010:
		return "ILLEGAL" //-lxOp0010" //ILLEGAL
		break
	case ClasslxOP0011:
		return "ILLEGAL" //-lxOP0011" //ILLEGAL
		break
	case Classlxop0100:
		return "ILLEGAL" //-lXop0100" //ILLEGAL
		break
	case ClasslxoP0101:
		return "ILLEGAL" //-lXoP0101" //ILLEGAL
		break
	case ClasslxOp0110:
		return "ILLEGAL" //-lXOp0110" //ILLEGAL
	case ClasslxOP0111:
		return "ILLEGAL" //-lXOP0111" //ILLEGAL
		break
	case Classlxop1000:
		return "DRAW" //-Lxop1000" //DRAW  Legal board Neither O or X win and game is full
		break
	case ClasslxoP1001:
		return "PLAY" //-LxoP1001" PLAY  Legal board Neither O or X win and game is in play
		break
	case ClasslxOp1010:
		return "OWIN" //-LxOp1010" OWIN  Legal board OWins and game is full
		break
	case ClasslxOP1011:
		return "OWIN" //-LxOP1011" OWIN  Legal board OWins and game is in play
		break
	case Classlxop1100:
		return "XWIN" //-LXop1100" //XWIN  Legal board XWins and game is full
		break
	case ClasslxoP1101:
		return "XWIN" //-LXoP1101" //XWIN  Legal board XWins and game is in play
		break
	case ClasslxOp1110:
		return "ERROR" //-LXOp1110" //ERROR Legal board but has both X AND O winning lines and is full
		//In normal play one win will be detected for either O or X, there will be both!
		break
	case ClasslxOP1111:
		return "ERROR" //-LXOP1111" //ERROR Legal board has both X and O winning but is in play
		break
	default:
		return "UNCLASSIFIED"

	}
	return ""
}

/*
switch GameStateclass(gs) {
case Classlxop0000:
	gsl[s] = "ILLEGAL" //-lxop0000"
	break
case ClasslxoP0001:
	gsl[s] = "ILLEGAL" //-lxoP0001" //ILLEGAL
	break
case ClasslxOp0010:
	gsl[s] = "ILLEGAL" //-lxOp0010" //ILLEGAL
	break
case ClasslxOP0011:
	gsl[s] = "ILLEGAL" //-lxOP0011" //ILLEGAL
	break
case Classlxop0100:
	gsl[s] = "ILLEGAL" //-lXop0100" //ILLEGAL
	break
case ClasslxoP0101:
	gsl[s] = "ILLEGAL" //-lXoP0101" //ILLEGAL
	break
case ClasslxOp0110:
	gsl[s] = "ILLEGAL" //-lXOp0110" //ILLEGAL
case ClasslxOP0111:
	gsl[s] = "ILLEGAL" //-lXOP0111" //ILLEGAL
	break
case Classlxop1000:
	gsl[s] = "DRAW" //-Lxop1000" //DRAW  Legal board Neither O or X win and game is full
	break
case ClasslxoP1001:
	gsl[s] = "PLAY" //-LxoP1001" PLAY  Legal board Neither O or X win and game is in play
	break
case ClasslxOp1010:
	gsl[s] = "OWIN" //-LxOp1010" OWIN  Legal board OWins and game is full
	break
case ClasslxOP1011:
	gsl[s] = "OWIN" //-LxOP1011" OWIN  Legal board OWins and game is in play
	break
case Classlxop1100:
	gsl[s] = "XWIN" //-LXop1100" //XWIN  Legal board XWins and game is full
	break
case ClasslxoP1101:
	gsl[s] = "XWIN" //-LXoP1101" //XWIN  Legal board XWins and game is in play
	break
case ClasslxOp1110:
	gsl[s] = "ERROR" //-LXOp1110" //ERROR Legal board but has both X AND O winning lines and is full
	//In normal play one win will be detected for either O or X, there will be both!
	break
case ClasslxOP1111:
	gsl[s] = "ERROR" //-LXOP1111" //ERROR Legal board has both X and O winning but is in play
	break
default:
	def++

} */
