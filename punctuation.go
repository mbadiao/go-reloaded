package main

import "strings"

func punc(file []string) []string {
	var final string
	text := strings.Join(file, " ")
	tab := []rune(text)
	for i := 1; i < len(tab); i++ {
		if tab[i] == '.' || tab[i] == ',' || tab[i] == '!' || tab[i] == '?' || tab[i] == ':' || tab[i] == ';' {
			for tab[i-1] == ' ' {
				tab = append(tab[:i-1], tab[i:]...)
			}
		}
	}
	for i := 0; i < len(tab); i++ {
		final += string(tab[i])
		if tab[i] == '.' || tab[i] == ',' || tab[i] == '!' || tab[i] == '?' || tab[i] == ':' || tab[i] == ';' {
			if i < len(tab)-1 {
				if tab[i+1] != '.' && tab[i+1] != '!' && tab[i+1] != '?' && tab[i+1] != ':' && tab[i+1] != ',' && tab[i+1] != ';' {
					final += " "
				}
			}
		}
	}
	return strings.Fields(final)
}
