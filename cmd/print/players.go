/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package print

import (
	"fmt"

	"github.com/spf13/cobra"
)

// playersCmd represents the players command
var playersCmd = &cobra.Command{
	Use:   "players",
	Short: "print players",
	Long: `All about players`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("players called")
	},
}

func init() {
	PrintCmd.AddCommand(playersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
