package main

import (
	"strings"
	"unicode"
)

func aps(tabWord []string) []string {
	var count int
	var text string
	var result []string
	str := strings.Join(tabWord, " ")
	// Gestion du nombre de ponctuation impaire
	for i := 0; i < len(str); i++ {
		if i > 0 && i < len(str)-1 {
			if str[i] == '\'' && (!unicode.IsLetter(rune(str[i-1])) || !unicode.IsLetter(rune(str[i+1]))) {
				count++
			}
		} else if i == 0 && str[i] == '\'' && i < len(str)-1 {
			count++
		} else if i == len(str)-1 && str[i] == '\'' && i > 0 {
			count++
		}
	}
	if count%2 != 0 {
		return []string{"Nombre", "quotes", "impair"}
	}
	count = 0
	// Bonne Remise des ponctuations
	for i := 0; i < len(str); i++ {
		text += string(str[i])
		if i > 0 && i < len(str)-1 {
			if str[i] == '\'' && (!unicode.IsLetter(rune(str[i-1])) || !unicode.IsLetter(rune(str[i+1]))) {
				count++
				if count%2 != 0 {
					if str[i-1] != 32 {
						text = text[:len(text)-1] + " " + text[len(text)-1:]
					}
					if str[i+1] == 32 {
						i++
					}
				} else {
					if str[i-1] == ' ' && text[len(text)-2] == ' ' {
						text = text[:len(text)-2]
						text += "'"
					}
					if str[i+1] != ' ' {
						text += " "
					}
				}
			}
		} else if i == 0 && str[i] == '\'' && i < len(str)-1 {
			count++
			if str[i+1] == ' ' {
				i++
			}
		} else if i == len(str)-1 && str[i] == '\'' && i > 0 {
			count++
			if str[i-1] == ' ' && text[len(text)-2] == ' ' {
				text = text[:len(text)-2]
				text += "'"
			}
		}
	}
	result = strings.Fields(text)
	return result
}
