package main

import (
	"strconv"
	"strings"
	"unicode"
)

func hexbin(tabWord []string) []string {
	for i := 1; i < len(tabWord); i++ {
		var surplus string
		if strings.Contains(tabWord[i], "(hex)") {
			// Gestion de la transformation
			nbChar := len(tabWord[i-1])
			for k := 0; k < len(tabWord[i-1]); k++ {
				if unicode.IsNumber(rune(tabWord[i-1][k])) || unicode.IsLetter(rune(tabWord[i-1][k])) {
					nbChar -= 1
				}
			}
			hex, _ := strconv.ParseInt(tabWord[i-1][nbChar:], 16, 64)
			if strconv.Itoa(int(hex)) != "0" {
				tabWord[i-1] = strconv.Itoa(int(hex))
			}
			//Gestion des Surplus
			for l := 5; l < len(tabWord[i]); l++ {
				surplus += tabWord[i][l:]
			}
			tabWord[i-1] += surplus
			tabWord = append(tabWord[:i], tabWord[i+1:]...)
			i--
		} else if strings.Contains(tabWord[i], "(bin)") {
			// Gestion de la transformation
			nbChar := len(tabWord[i-1])
			for k := 0; k < len(tabWord[i-1]); k++ {
				if unicode.IsNumber(rune(tabWord[i-1][k])) || unicode.IsLetter(rune(tabWord[i-1][k])) {
					nbChar -= 1
				}
			}
			bin, _ := strconv.ParseInt(tabWord[i-1][nbChar:], 2, 64)
			if strconv.Itoa(int(bin)) != "0" {
				tabWord[i-1] = strconv.Itoa(int(bin))
			}
			//Surplus
			for l := 5; l < len(tabWord[i]); l++ {
				surplus += tabWord[i][l:]
			}
			tabWord[i-1] += surplus
			tabWord = append(tabWord[:i], tabWord[i+1:]...)
			i--
		}
	}
	return tabWord
}
