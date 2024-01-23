/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package file

import (
	"fmt"

	"github.com/spf13/cobra"
)

// directoryCmd represents the directory command
var directoryCmd = &cobra.Command{
	Use:   "directory",
	Short: "please choose a directory for files",
	Long: `All about directories`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("directory called")
	},
}

func init() {
	FileCmd.AddCommand(directoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// directoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// directoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
