package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gorecords",
	Short: "A tool for reading records from file",
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(serverCmd)

}
