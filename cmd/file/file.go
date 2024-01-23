/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package file

import (
//	"fmt"

	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var FileCmd = &cobra.Command{
	Use:   "file",
	Short: "File commands",
	Long: `file commands : save, load, list, delete, directory, info`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
//	rootCmd.AddCommand(fileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
