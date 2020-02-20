package passwordGenerator

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

type Settings struct {
	Length                  int
	IncludeUpperCaseLetters bool
	IncludeLowerCaseLetters bool
	IncludeDigits           bool
	IncludeSymbols          bool
}

func ValidateSettings(settings *Settings) (err error) {
	if settings.Length == 0 {
		err = fmt.Errorf("Password length can't be zero")
	}
	if !(settings.IncludeSymbols || settings.IncludeLowerCaseLetters || settings.IncludeDigits || settings.IncludeUpperCaseLetters) {
		err = fmt.Errorf("At least one criterion has to be set")
	}
	return
}

func getValidChars(settings *Settings) (validChars []byte) {
	var c byte
	for c = 0; c < 128; c++ {
		if unicode.IsLower(rune(c)) && settings.IncludeLowerCaseLetters ||
			unicode.IsUpper(rune(c)) && settings.IncludeUpperCaseLetters ||
			unicode.IsDigit(rune(c)) && settings.IncludeDigits ||
			unicode.IsSymbol(rune(c)) && settings.IncludeSymbols {
			validChars = append(validChars, c)
		}
	}
	return
}

func GeneratePassword(settings *Settings) []rune {
	password := make([]rune, settings.Length)
	rand.Seed(time.Now().UTC().UnixNano())
	validChars := getValidChars(settings)
	for i := 0; i < settings.Length; i++ {
		randomIndex := rand.Intn(len(validChars))
		password[i] = rune(validChars[randomIndex])
	}
	return password
}

