package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "passGen",
	Short: "PassGen",
	Long:  "Password Generator",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("started")
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
