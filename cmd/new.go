package cmd

import (
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/rjNemo/go-pass-gen/passGen"
)

func init() {
	newPasswordCommand.Flags().IntVarP(&Length, "length", "l", 6, "password length")
	newPasswordCommand.Flags().BoolVarP(&WithNumbers, "numbers", "n", false, "password should contain numbers")
	rootCommand.AddCommand(newPasswordCommand)
}

// newPasswordCommand creates a new password. It takes into account the supplied flags.
var newPasswordCommand = &cobra.Command{Use: "new",
	Short: "New Password",
	Long:  "Create a secure password",
	Run: func(cmd *cobra.Command, args []string) {
		pg := passGen.NewPasswordGenerator(passGen.Options{Length: Length, WithNumbers: WithNumbers})
		password := pg.NewPassword()
		display(password)
	}}

// display presents information to standard output.
func display(str string) {
	red := color.New(color.FgGreen).Add(color.Bold)
	whiteBackground := red.Add(color.BgWhite)
	_, err := whiteBackground.Println(str)
	if err != nil {
		log.Fatal(err)
		return
	}
}
