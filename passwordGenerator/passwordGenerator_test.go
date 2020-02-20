package passwordGenerator

import (
	"fmt"
	"testing"
	"unicode"
)

func TestValidateSettings(t *testing.T) {
	tables := []struct {
		s       Settings
		isError bool
	}{
		{
			s:       Settings{Length: 5, IncludeUpperCaseLetters: true, IncludeLowerCaseLetters: true, IncludeDigits: true, IncludeSymbols: false},
			isError: false,
		},
		{
			s:       Settings{Length: 0, IncludeUpperCaseLetters: true, IncludeLowerCaseLetters: true, IncludeDigits: true, IncludeSymbols: false},
			isError: true,
		},
		{
			s:       Settings{Length: 5, IncludeUpperCaseLetters: false, IncludeLowerCaseLetters: false, IncludeDigits: false, IncludeSymbols: false},
			isError: true,
		},
	}

	for _, table := range tables {
		err := ValidateSettings(&table.s)
		if err != nil && !table.isError {
			t.Errorf("Error was not expected for %v", table.s)
		}
		if err == nil && table.isError {
			t.Errorf("Error was expected for %v", table.s)
		}
	}
}

func TestGetValidChars(t *testing.T) {
	tables := []struct {
		s       Settings
	}{
		{
			s:       Settings{Length: 5, IncludeUpperCaseLetters: true, IncludeLowerCaseLetters: true, IncludeDigits: true, IncludeSymbols: false},
		},
		{
			s:       Settings{Length: 0, IncludeUpperCaseLetters: true, IncludeLowerCaseLetters: true, IncludeDigits: true, IncludeSymbols: false},
		},
		{
			s:       Settings{Length: 5, IncludeUpperCaseLetters: true, IncludeLowerCaseLetters: false, IncludeDigits: false, IncludeSymbols: false},
		},
		{
			s:       Settings{Length: 5, IncludeUpperCaseLetters: false, IncludeLowerCaseLetters: true, IncludeDigits: false, IncludeSymbols: true},
		},
	}

	for _, table := range tables {
		validChars := getValidChars(&table.s)
		validCharsR := make([]rune,len(validChars))	
		for i, c := range validChars {
			validCharsR[i] = rune(c)
		}
		testIfCharsAreValid(t, validCharsR, &table.s)
	}
}

func TestGeneratePassword(t *testing.T) {
	tables := []struct {
		s       Settings
	}{
		{
			s:       Settings{Length: 5, IncludeUpperCaseLetters: true, IncludeLowerCaseLetters: true, IncludeDigits: true, IncludeSymbols: false},
		},
		{
			s:       Settings{Length: 4, IncludeUpperCaseLetters: true, IncludeLowerCaseLetters: true, IncludeDigits: true, IncludeSymbols: false},
		},
		{
			s:       Settings{Length: 5, IncludeUpperCaseLetters: false, IncludeLowerCaseLetters: true, IncludeDigits: false, IncludeSymbols: false},
		},
		{
			s:       Settings{Length: 5, IncludeUpperCaseLetters: false, IncludeLowerCaseLetters: true, IncludeDigits: false, IncludeSymbols: false},
		},
	}

	for _, table := range tables {
		password := GeneratePassword(&table.s)
		testIfCharsAreValid(t, password, &table.s)
	}
}

func testIfCharsAreValid(t *testing.T, chars []rune, s *Settings) {
	for _, c := range chars {
		if !s.IncludeLowerCaseLetters && unicode.IsLower(c) {
			fmt.Println("Should fail")
			t.Errorf("Lower case letter not allowed")
		}
		if !s.IncludeUpperCaseLetters && unicode.IsUpper(c) {
			t.Errorf("Upper case letter not allowed")
		}
		if !s.IncludeDigits && unicode.IsNumber(c) {
			t.Errorf("Digits not allowed")
		}
		if !s.IncludeSymbols && unicode.IsSymbol(c) {
			t.Errorf("Symbol not allowed")
		}
	}
}
