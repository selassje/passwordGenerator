package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

type Settings struct {
	Length                  int
	IncludeUpperCaseLetters bool
	IncludeLowerCaseLetter  bool
	IncludeDigits           bool
	IncludeSymbols          bool
}

func validateSettings(settings *Settings) (err error) {
	if settings.Length == 0 {
		err = fmt.Errorf("Password length can't be zero")
	}
	if !(settings.IncludeSymbols || settings.IncludeLowerCaseLetter || settings.IncludeDigits || settings.IncludeUpperCaseLetters) {
		err = fmt.Errorf("At least one criterion has to be set")
	}
	return
}

func getValidChars(settings *Settings) (validChars []byte) {
	var c byte
	for c = 0; c < 128; c++ {
		if unicode.IsLower(rune(c)) && settings.IncludeLowerCaseLetter ||
			unicode.IsUpper(rune(c)) && settings.IncludeUpperCaseLetters ||
			unicode.IsDigit(rune(c)) && settings.IncludeDigits ||
			unicode.IsSymbol(rune(c)) && settings.IncludeSymbols {
			validChars = append(validChars, c)
		}
	}
	return
}

func generatePassword(settings *Settings) []rune {
	password := make([]rune, settings.Length)
	rand.Seed(time.Now().UTC().UnixNano())
	validChars := getValidChars(settings)
	fmt.Println(validChars)
	for i := 0; i < settings.Length; i++ {
		randomIndex := rand.Intn(len(validChars))
		password[i] = rune(validChars[randomIndex])
	}
	return password
}

func main() {
	settings := Settings{
		Length:                  32,
		IncludeUpperCaseLetters: true,
		IncludeLowerCaseLetter:  true,
		IncludeDigits:           true,
		IncludeSymbols:          false,
	}
	err := validateSettings(&settings)
	if err != nil {
		panic(err.Error())
	}
	password := generatePassword(&settings)
	fmt.Println(string(password))
}
