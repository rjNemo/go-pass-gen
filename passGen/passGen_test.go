package passGen_test

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/rjNemo/go-pass-gen/passGen"
)

func TestGeneratePasswordWithGivenCharacterNumber(t *testing.T) {
	opts := passGen.Options{
		Length: rand.Intn(12),
	}
	pg := passGen.NewPasswordGenerator(opts)

	if password := pg.NewPassword(); len(password) != opts.Length {
		t.Errorf("Expected a password to be %d characters long, got %d", opts.Length, len(password))
	}
}

func TestGeneratePasswordWithDefaultCharacterNumber(t *testing.T) {
	pg := passGen.NewPasswordGenerator(passGen.Options{})

	if password := pg.NewPassword(); len(password) != 6 {
		t.Errorf("Expected a password to be %d characters long, got %d", 6, len(password))
	}
}

func TestGeneratePasswordWithLettersAndNumbers(t *testing.T) {
	opts := passGen.Options{
		WithNumbers: true,
	}
	pg := passGen.NewPasswordGenerator(opts)

	if password := pg.NewPassword(); !containNumbers(password) {
		t.Errorf("Expected password to contain NUMBERS, got %q", password)
	}
}

func containNumbers(str string) bool {
	return strings.ContainsAny(str, passGen.NUMBERS)
}
