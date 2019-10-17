package service

import "strings"

var nameLetters map[rune]bool

func init() {
	nameLetters = make(map[rune]bool)
	nameLetters['_'] = true
	for r := 'A'; r <= 'Z'; r++ {
		nameLetters[r] = true
	}
	for r := 'a'; r <= 'z'; r++ {
		nameLetters[r] = true
	}
	for r := '0'; r <= '9'; r++ {
		nameLetters[r] = true
	}
}

func f(r rune) bool {
	return !nameLetters[r]
}

//ValidateName 校验名称合法性
func ValidateName(name string) bool {
	if name == "" {
		return false
	}
	return strings.IndexFunc(name, f) == -1

}
