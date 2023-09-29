package main

import "strings"

func punc(tabWord []string) []string {
	var final string
	text := strings.Join(tabWord, " ")
	tabChar := []rune(text)
	// Retrait space arrière
	for i := 1; i < len(tabChar); i++ {
		if tabChar[i] == '.' || tabChar[i] == ',' || tabChar[i] == ';' || tabChar[i] == '!' || tabChar[i] == '?' || tabChar[i] == ':' || tabChar[i] == '|' || tabChar[i] == '(' || tabChar[i] == ')' || tabChar[i] == 39 {
			for tabChar[i-1] == ' ' {
				tabChar = append(tabChar[:i-1], tabChar[i:]...)
			}
		}
	}
	// Ajout space avant
	for i := 0; i < len(tabChar); i++ {
		final += string(tabChar[i])
		if tabChar[i] == '.' || tabChar[i] == ',' || tabChar[i] == ';' || tabChar[i] == '!' || tabChar[i] == '?' || tabChar[i] == ':' {
			if i < len(tabChar)-1 {
				if tabChar[i+1] != '.' && tabChar[i+1] != ',' && tabChar[i+1] != ';' && tabChar[i+1] != '!' && tabChar[i+1] != '?' && tabChar[i+1] != ':' && tabChar[i+1] != 39 {
					final += " "
				}
			}
		}
	}
	result := strings.Fields(final)
	return result
}
