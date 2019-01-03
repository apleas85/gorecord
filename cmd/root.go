package cmd

import (
	"github.com/spf13/cobra"
)

var files string
var rootCmd = &cobra.Command{
	Use:   "gorecords",
	Short: "A tool for reading records from file",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&files, "files", "spaceRecords.txt,commaRecords.txt,pipeRecords.txt", "Specify 3 files seperated by ','")
}
func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(serverCmd)

}
