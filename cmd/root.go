// Package cmd defines the command-line interface to passgen
package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

// rootCommand initializes the command-line interface application.
var rootCommand = &cobra.Command{
	Use:   "passgen",
	Short: "PassGen",
	Long:  "Password Generator",
	Run: func(cmd *cobra.Command, args []string) {
		display("** passgen – v0.0.1 **\nUse passgen -h for more information")
	},
}

// Execute is where the fun actually happens.
func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
