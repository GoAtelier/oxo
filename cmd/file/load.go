/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package file

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load a file",
	Long: `All about file loading`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("load called")
	},
}

func init() {
	FileCmd.AddCommand(loadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
