package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

// rootCommand initializes the command-line interface application.
var rootCommand = &cobra.Command{
	Use:   "passGen",
	Short: "PassGen",
	Long:  "Password Generator",
	Run: func(cmd *cobra.Command, args []string) {
		display("** passGen – v0.0.1 **\nUse passGen -h for more information")
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
