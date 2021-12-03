package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "api",
	Short: "Run API server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting API server")
	},
}

func Run() {
	rootCommand.AddCommand(serverCommand)

	if err := rootCommand.Execute(); err != nil {
		log.Panic(err)
	}
}
