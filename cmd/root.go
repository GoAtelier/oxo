/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/finecoding/oxo-1/oxo"
	"github.com/spf13/cobra"
)

var numgames int //Command line flags will be stored in these variables
var oplayer, xplayer string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "oxocli",

	Short: "Utility to experiment with a simulation of the game of Noughts and Crosses (aka Tic Tac Toe)",
	Long: `This accepts flags for the 
	number of games to be played Example: -n 10, 
	tactic used by O Example: -o RANDOM
	tactic used by X Example: -x RANDOM
	The tactics available are RANDOM, CENTRE, CORNER, COMPLETEXLINE and COMPLETEOLINE
	Example:
	oxocli -n 10 -o RANDOM -x CENTRE
	The default values are -n 1 -o RANDOM -x RANDOM`,
	Run: flagsFunc,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func flagsFunc(cmd *cobra.Command, args []string) {
	nl := oxo.Newlookup()										//The state looktable that maps a string representation of a board with its state

	var res oxo.Group											//Group is a slice of Games, it will hold all the Game records
	var flipwhostarts bool									

	playerlookup := oxo.NewPlayerlookup()						//playerlookup is a map of strings to predefined Players.

	for i := 0; i < numgames; i++ {									//repeats the number of times specified by the command line -n	
		flipwhostarts = !flipwhostarts                         	//alternates who starts the game eliminating any first mover advantage
		OPlayer := playerlookup[oplayer]                       	//Setup O. Lookup the string entered at the command line following -o to find the Player
		XPlayer := playerlookup[xplayer]                       	//Setup X. Lookup the string entered at the command line following -x to find the player
		x := oxo.Playgame(nl, OPlayer, XPlayer, flipwhostarts) 	//PlayGame, with two Players, they will take turns starting each game to eliminate any advantage
		res.Games = append(res.Games, x)                       	//The completed Game is added to the Games slice.
		res.UpdateNums(x)									   	//Update some basic statistics about each game  
		fmt.Println(res.Games[i])							   	//Print the game across the screen.....we need a flag for this!

	}

	fmt.Printf("Printed %d games\n\n", numgames)					//Print some stats about all the games played.
	fmt.Printf(" Number of XWINS = %d\n", res.NumXwins)
	fmt.Printf(" Number of OWINS = %d\n", res.NumOwins)
	fmt.Printf(" Number of DRAWS = %d\n", res.NumDraws)
	fmt.Printf(" Number of ILLEGALS = %d\n", res.NumIllegals)

}
func init() {
	//Obtaining flags from the Command line.
	rootCmd.Flags().IntVarP(&numgames, "numgames", "n", 1, "number of games to play")
	rootCmd.Flags().StringVarP(&oplayer, "oplayer", "o", "RANDOM", "O player,  select a tactic RANDOM, CENTRE. CORNER....")
	rootCmd.Flags().StringVarP(&xplayer, "xplayer", "x", "RANDOM", "X player. select a tactic RANDOM, CENTRE. CORNER....")
}
