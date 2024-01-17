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

var reps int
var oplayer, xplayer string
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "oxocli",

	Short: "Utility to experiment with a simulation of the game of Noughts and Crosses (aka Tic Tac Toe)",
	Long:  `Enter two players and the number of games to play`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
//	fmt.Printf("p = %d\n", reps)
//	fmt.Println("args:", args)
	nl := oxo.Newlookup()

//Now to add these to a new data structure, a Result, which will be a slice of Games.

var res oxo.Group
var flipwhostarts bool

playerlookup:=oxo.NewPlayerlookup()


for i := 0; i < reps; i++ {
	flipwhostarts = !flipwhostarts           //alternates who starts the game eliminating any first mover advantage
	OPlayer := playerlookup[oplayer]        //Setup O looksup the string entered at the command line for -o to finds the player 
	XPlayer := playerlookup[xplayer]        //Setup X looksup the string entered at the command line for -x to finds the player 
	x := oxo.Playgame(nl, OPlayer, XPlayer, flipwhostarts) //PlayGame, toss coin to decide who goes first...
	res.Games = append(res.Games, x)                       //Add Game to res slice.
	res.UpdateNums(x)
	fmt.Println(res.Games[i]) 

}	

fmt.Printf("Printed %d games\n\n",reps)
fmt.Printf(" Number of XWINS = %d\n", res.NumXwins)
fmt.Printf(" Number of OWINS = %d\n", res.NumOwins)
fmt.Printf(" Number of DRAWS = %d\n", res.NumDraws)
fmt.Printf(" Number of ILLEGALS = %d\n", res.NumIllegals)

}
func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.oxo-1.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().IntVarP(&reps, "numgames", "n", 1, "number of games to play")
	rootCmd.Flags().StringVarP(&oplayer, "oplayer", "o", "RANDOM", "O player,  select a tactic RANDOM, CENTRE. CORNER....")
	rootCmd.Flags().StringVarP(&xplayer, "xplayer", "x", "RANDOM", "X player. select a tactic RANDOM, CENTRE. CORNER....")
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
