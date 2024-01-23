/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package print

import (
	"fmt"

	"github.com/spf13/cobra"
)

// printCmd represents the print command
var PrintCmd = &cobra.Command{
	Use:   "print",
	Short: "print stuff",
	Long: `All about printing`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("print called")
	},
}

func init() {
//	rootCmd.AddCommand(PrintCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
