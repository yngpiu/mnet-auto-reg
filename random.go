package main

import (
	"math/rand/v2"
	"strings"
)

type Birthdate struct {
	Year  int
	Month int
	Day   int
}

func randomInt(min, max int) int {
	return min + rand.IntN(max-min+1)
}

func randomBirthdate() Birthdate {
	return Birthdate{
		Year:  randomInt(1990, 2005),
		Month: randomInt(1, 12),
		Day:   randomInt(1, 28),
	}
}

func randomGender() string {
	if rand.IntN(2) == 0 {
		return "f"
	}
	return "m"
}

type KoreanName struct {
	Surname string
	Given   string
	Full    string
}

func randomKoreanName() KoreanName {
	surname := koreanSurnames[rand.IntN(len(koreanSurnames))]
	given := koreanGivenNames[rand.IntN(len(koreanGivenNames))]
	return KoreanName{
		Surname: surname,
		Given:   given,
		Full:    surname + given,
	}
}

func validatePassword(pwd string) (bool, []string) {
	var errors []string
	if len(pwd) < 8 || len(pwd) > 20 {
		errors = append(errors, "length (8-20 characters)")
	}
	if !containsAny(pwd, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		errors = append(errors, "uppercase letter")
	}
	if !containsAny(pwd, "abcdefghijklmnopqrstuvwxyz") {
		errors = append(errors, "lowercase letter")
	}
	if !containsAny(pwd, "0123456789") {
		errors = append(errors, "number")
	}
	if !containsAny(pwd, "!@#$%^&*()_+-=[]{}|;':\",./<>?~`") {
		errors = append(errors, "special character")
	}
	return len(errors) == 0, errors
}

func containsAny(s, chars string) bool {
	for _, c := range chars {
		if strings.ContainsRune(s, c) {
			return true
		}
	}
	return false
}
