/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package file

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("save called")

		_, err := os.Stat(filename)
		if os.IsNotExist(err) {
			file, err := os.Create(filename)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
		} else {
			fmt.Println("File already exists!", filename)
			return
		}

		fmt.Println("File created successfully:", filename)


		
	},
}
var filename string

func init() {
	FileCmd.AddCommand(saveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	saveCmd.Flags().StringVarP(&filename, "filename", "f", "default.txt", "Please enter a file name")

}
