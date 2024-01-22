/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/finecoding/oxo-1/oxo"
	"github.com/spf13/cobra"
)

var numgames int //Command line flags will be stored in these variables
var oplayer, xplayer string
var printonscreen, printjson bool

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
	RunE: flagsFunc,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func validateFlags(n int, op string, xp string) error {
	var listoftactics = []string{"RANDOM", "CENTRE", "CORNER", "COMPLETEXLINE", "COMPLETOLINE"}
	if !slices.Contains(listoftactics, op) {
		return errors.New("Unrecognised tactic for o. Enter one of RANDOM,CENTRE,CORNER,COMPLETEXLINE,COMPLETOLINE")
	}
	if !slices.Contains(listoftactics, xp) {
		return errors.New("Unrecognised tactic for x. Enter one of RANDOM,CENTRE,CORNER,COMPLETEXLINE,COMPLETOLINE")
	}
	if n < 0 {
		return errors.New("n cannot be smaller that 0")
	}

	if op == "" {
		return errors.New("enter a tactic name for o")
	}
	if xp == "" {
		return errors.New("enter a tactic name for x")
	}
	return nil
}
func flagsFunc(cmd *cobra.Command, args []string) error {
	logfile, err := os.Create("app.log")

	if err != nil {
		log.Fatal(err)
	}

	defer logfile.Close()
	log.SetOutput(logfile)
	log.Println("Starting the application...")

	var res oxo.Group //Group is a slice of Games, it will hold all the Game records
	var flipwhostarts bool
	err = validateFlags(numgames, oplayer, xplayer)
	if err != nil {
		return err
	}

	playerlookup := oxo.NewPlayerlookup() //playerlookup is a map of strings to predefined Players.
	boardstatelookup := oxo.Newlookup()   //The board state lookup table that maps a string representation of a board with its state
	for i := 0; i < numgames; i++ {       //repeats the number of times specified by the command line -n
		log.Println("Info Starting Game Loop")
		flipwhostarts = !flipwhostarts                                       //alternates who starts the game eliminating any first mover advantage
		OPlayer := playerlookup[oplayer]                                     //Setup O. Lookup the string entered at the command line following -o to find the Player
		XPlayer := playerlookup[xplayer]                                     //Setup X. Lookup the string entered at the command line following -x to find the player
		x := oxo.Playgame(boardstatelookup, OPlayer, XPlayer, flipwhostarts) //PlayGame, with two Players, they will take turns starting each game to eliminate any advantage
		res.Games = append(res.Games, x)                                     //The completed Game is added to the Games slice.
		res.UpdateNums(x)                                                    //Update some basic statistics about each game
		if printonscreen {
			fmt.Println(res.Games[i])
		} //Print the games on the terminal is off by default, prints only stats. -p or --printonscreen flag switches on printing all games.
		log.Println("Info End Game Loop")
	}

	fmt.Printf("Printed %d games\n\n", numgames) //Print some stats about all the games played.
	fmt.Printf(" Number of XWINS = %d\n", res.NumXwins)
	fmt.Printf(" Number of OWINS = %d\n", res.NumOwins)
	fmt.Printf(" Number of DRAWS = %d\n", res.NumDraws)
	fmt.Printf(" Number of ILLEGALS = %d\n", res.NumIllegals) //This should always be zero!
	if printjson {
		bs, err := json.MarshalIndent(res, "", " ") //This will print a lot of json for each game! So keep numgames low
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(bs))
	}
	return nil
}
func init() {
	//Obtaining flags from the Command line.
	rootCmd.Flags().BoolVarP(&printonscreen, "printonscreen", "p", false, "print games on terminal")
	rootCmd.Flags().BoolVarP(&printjson, "printjson", "j", false, "print json on terminal")
	rootCmd.Flags().IntVarP(&numgames, "numgames", "n", 1, "number of games to play")
	rootCmd.Flags().StringVarP(&oplayer, "oplayer", "o", "RANDOM", "O player,  select a tactic RANDOM, CENTRE, CORNER,COMPLETEXLINE,COMPLETOLINE")
	rootCmd.Flags().StringVarP(&xplayer, "xplayer", "x", "RANDOM", "X player. select a tactic RANDOM, CENTRE, CORNER, COMPLETEXLINE,COMPLETOLINE")
}
