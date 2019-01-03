package cmd

import (
	"github.com/go/gorecords/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts server",
	Run:   startServerCmd,
}

func startServerCmd(*cobra.Command, []string) {
	server.Start(files)
}
