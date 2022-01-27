// Package passgen creates a random password.
package passgen

import (
	"math/rand"
)

// PasswordGenerator handles passwords creation.
type PasswordGenerator struct {
	characters []rune
	options    *Options
}

// New returns a valid PasswordGenerator given the specified Options.
func New(opts *Options) *PasswordGenerator {
	pg := &PasswordGenerator{}
	pg.options = opts.SetDefaults()
	pg.characters = pg.shuffleCharacters(opts.WithNumbers)
	return pg
}

// NewPassword returns a pseudo random string.
func (p PasswordGenerator) NewPassword() string {
	return p.generatePassword(p.options.Length)
}

// generatePassword builds the new password.
func (p PasswordGenerator) generatePassword(length int) string {
	var res string
	for i := 0; i < length; i++ {
		res += string(p.characters[i])
	}
	return res
}

// shuffleCharacters randomizes the characters.
func (p PasswordGenerator) shuffleCharacters(withNumbers bool) []rune {
	letters := []rune(UPPERCASE + LOWERCASE)
	if withNumbers {
		letters = append(letters, []rune(NUMBERS)...)
	}
	rand.Shuffle(len(letters), func(i, j int) { letters[i], letters[j] = letters[j], letters[i] })
	return letters
}
