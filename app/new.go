package app

import (
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/rjNemo/go-pass-gen/passgen"
)

var (
	// length of the generated password.
	length int
	// withNumbers is set to true if the new password must contain numbers.
	withNumbers bool
)

func init() {
	newPasswordCommand.Flags().IntVarP(&length, "length", "l", 6, "password length")
	newPasswordCommand.Flags().BoolVarP(&withNumbers, "numbers", "n", false, "password should contain numbers")
	rootCommand.AddCommand(newPasswordCommand)
}

// newPasswordCommand creates a new password. It takes into account the supplied flags.
var newPasswordCommand = &cobra.Command{Use: "new",
	Short: "New Password",
	Long:  "Create a secure password",
	Run: func(cmd *cobra.Command, args []string) {
		opts := passgen.Options{Length: length, WithNumbers: withNumbers}
		pg := passgen.New(opts.SetDefaults())
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
