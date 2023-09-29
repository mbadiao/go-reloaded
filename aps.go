package main

import (
	"strings"
	"unicode"
)

func aps(tabWord []string) []string {
	var cpt int
	var text string
	var result []string
	stringF := strings.Join(tabWord, " ")
	// Gestion du nombre de ponctuation impaire
	for i := 0; i < len(stringF); i++ {
		if i > 0 && i < len(stringF)-1 {
			if stringF[i] == '\'' && (!unicode.IsLetter(rune(stringF[i-1])) || !unicode.IsLetter(rune(stringF[i+1]))) {
				cpt++
			}
		} else if i == 0 && stringF[i] == '\'' && i < len(stringF)-1 {
			cpt++
		} else if i == len(stringF)-1 && stringF[i] == '\'' && i > 0 {
			cpt++
		}
	}
	if cpt%2 != 0 {
		return tabWord//[]string{"Le","Nombre", "de", "quotes","que vous avez saisie est", "impair"}
	}
	cpt = 0
	// Bonne Remise des ponctuations
	for i := 0; i < len(stringF); i++ {
		text += string(stringF[i])
		if i > 0 && i < len(stringF)-1 {
			if stringF[i] == '\'' && (!unicode.IsLetter(rune(stringF[i-1])) || !unicode.IsLetter(rune(stringF[i+1]))) {
				cpt++
				if cpt%2 != 0 {
					if stringF[i-1] != 32 {
						text = text[:len(text)-1] + " " + text[len(text)-1:]
					}
					if stringF[i+1] == 32 {
						i++
					}
				} else {
					if stringF[i-1] == ' ' && text[len(text)-2] == ' ' {
						text = text[:len(text)-2]
						text += "'"
					}
					if stringF[i+1] != ' ' {
						text += " "
					}
				}
			}
		} else if i == 0 && stringF[i] == '\'' && i < len(stringF)-1 {
			cpt++
			if stringF[i+1] == ' ' {
				i++
			}
		} else if i == len(stringF)-1 && stringF[i] == '\'' && i > 0 {
			cpt++
			if stringF[i-1] == ' ' && text[len(text)-2] == ' ' {
				text = text[:len(text)-2]
				text += "'"
			}
		}
	}
	result = strings.Fields(text)
	return result
}
