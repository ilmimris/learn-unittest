package cmd

import (
	"github.com/ilmimris/learn-unittest/transport/rest"
	"github.com/spf13/cobra"
)

var serverCommand = &cobra.Command{
	Use:    "api",
	PreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		rest := rest.NewRest()

		go rest.Serve()
		rest.SignalCheck()
	},
	PostRun: func(cmd *cobra.Command, args []string) {},
}
