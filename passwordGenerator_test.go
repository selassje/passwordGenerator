package main

import (
	"testing"
	"unicode"
)

func TestValidateSettings(t *testing.T) {
	tables := []struct {
		s       Settings
		isError bool
	}{
		{
			s:       Settings{Length: 5, IncludeUpperCaseLetters: true, IncludeLowerCaseLetter: true, IncludeDigits: true, IncludeSymbols: false},
			isError: false,
		},
		{
			s:       Settings{Length: 0, IncludeUpperCaseLetters: true, IncludeLowerCaseLetter: true, IncludeDigits: true, IncludeSymbols: false},
			isError: true,
		},
		{
			s:       Settings{Length: 5, IncludeUpperCaseLetters: false, IncludeLowerCaseLetter: false, IncludeDigits: false, IncludeSymbols: false},
			isError: true,
		},
	}

	for _, table := range tables {
		err := validateSettings(&table.s)
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
			s:       Settings{Length: 5, IncludeUpperCaseLetters: true, IncludeLowerCaseLetter: true, IncludeDigits: true, IncludeSymbols: false},
		},
		{
			s:       Settings{Length: 0, IncludeUpperCaseLetters: true, IncludeLowerCaseLetter: true, IncludeDigits: true, IncludeSymbols: false},
		},
		{
			s:       Settings{Length: 5, IncludeUpperCaseLetters: false, IncludeLowerCaseLetter: false, IncludeDigits: false, IncludeSymbols: false},
		},
		{
			s:       Settings{Length: 5, IncludeUpperCaseLetters: false, IncludeLowerCaseLetter: true, IncludeDigits: false, IncludeSymbols: false},
		},
	}

	for _, table := range tables {
		func(t *testing.T, s *Settings) {
			validChars := getValidChars(s)
			for _, c := range validChars {
				r := rune(c)
				if !s.IncludeLowerCaseLetter && unicode.IsLower(r) {
					t.Errorf("Lower case letter not allowed")
				}
				if !s.IncludeUpperCaseLetters && unicode.IsUpper(r) {
					t.Errorf("Upper case letter not allowed")
				}
				if !s.IncludeDigits && unicode.IsNumber(r) {
					t.Errorf("Digits not allowed")
				}
				if !s.IncludeSymbols && unicode.IsSymbol(r) {
					t.Errorf("Symbol not allowed")
				}
			}
		}(t, &table.s)
	}
}
