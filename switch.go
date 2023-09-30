package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func SwitchFunc(tabWord []string) []string {
	var nbChar, confirm, y int
	var surplus string
	for i := 0; i < len(tabWord); i++ {
		nbChar, y = 0, 1
		surplus = ""
		if strings.Contains(tabWord[i], "(up)") {
			// Gestion de la transformation
			if i > 0 {
				for !isAlphaNum(tabWord[i-y]) {
					if i <= y {
						break
					}
					y++
				}
				tabWord[i-y] = strings.ToUpper(tabWord[i-y])
				//Surplus
				fmt.Println(surplus)
				for l := 5; l < len(tabWord[i]); l++ {
					surplus += tabWord[i][l:]
				}
				fmt.Println(surplus)
				tabWord[i-1] += surplus
			}
			tabWord = append(tabWord[:i], tabWord[i+1:]...)
			i--
		} else if strings.Contains(tabWord[i], "(low)") {
			// Gestion de la transformation
			if i > 0 {
				for !isAlphaNum(tabWord[i-y]) {
					if i <= y {
						break
					}
					y++
				}
				tabWord[i-y] = strings.ToLower(tabWord[i-y])
				//Surplus
				for l := 6; l < len(tabWord[i]); l++ {
					surplus += tabWord[i][l:]
				}
				tabWord[i-1] += surplus
			}
			tabWord = append(tabWord[:i], tabWord[i+1:]...)
			i--
		} else if strings.Contains(tabWord[i], "(cap)") {
			// Gestion de la transformation
			if i > 0 {
				for !isAlphaNum(tabWord[i-y]) {
					if i <= y {
						break
					}
					y++
				}
				tabWord[i-y] = strings.ToLower(tabWord[i-y])
				tabWord[i-y] = cap(tabWord[i-y])
				//Surplus
				for l := 6; l < len(tabWord[i]); l++ {
					surplus += tabWord[i][l:]
				}
				tabWord[i-1] += surplus
			}
			tabWord = append(tabWord[:i], tabWord[i+1:]...)
			i--
		} else if tabWord[i] == "(up," && i+1 < len(tabWord) {
			if i > 0 {
				// nbre de Caractère
				for k := 0; k < len(tabWord[i+1]); k++ {
					if tabWord[i+1][k] == '-' || tabWord[i+1][k] == '+' {
						continue
					}
					if !unicode.IsNumber(rune(tabWord[i+1][k])) {
						nbChar++
					}
					if k > 1 {
						if tabWord[i+1][k] == ')' && !unicode.IsNumber(rune(tabWord[i+1][k-1])) || !strings.Contains(tabWord[i+1], ")") {
							return []string{"Ce", "n'est", "pas", "une instance"}
						}
					}
				}
				// Gestion de la transformation
				nbre, _ := strconv.Atoi(tabWord[i+1][:len(tabWord[i+1])-nbChar])
				// fmt.Println(nbre)
				if nbre > i {
					nbre = i
				}
				for !isAlphaNum(tabWord[i-y]) {
					if i <= y {
						break
					}
					y++
				//	fmt.Println(y)
				}
				for j := y; j <= nbre+(y-1); j++ {
				//	fmt.Println(y,j)
					tabWord[i-j] = strings.ToUpper(tabWord[i-j])
				//	fmt.Println(tabWord[i-j])
				}
				confirm = nbre
			}
			// Suppression
			if nbChar != len(tabWord[i+1]) {
				tabWord = append(tabWord[:i], tabWord[i+2:]...)
			}
			if confirm != 0 {
				i--
			}
		} else if tabWord[i] == "(cap," && i+1 < len(tabWord) {
			if i > 0 {
				// nbre de Caractère
				for k := 0; k < len(tabWord[i+1]); k++ {
					if tabWord[i+1][k] == '-' || tabWord[i+1][k] == '+' {
						continue
					}
					if !unicode.IsNumber(rune(tabWord[i+1][k])) {
						nbChar++
					}
				}
				// Gestion de la transformation
				nbre, _ := strconv.Atoi(tabWord[i+1][:len(tabWord[i+1])-nbChar])
				if nbre > i {
					nbre = i
				}
				for !isAlphaNum(tabWord[i-y]) {
					if i <= y {
						break
					}
					y++
				}
				for j := y; j <= nbre+(y-1); j++ {
					tabWord[i-j] = strings.ToLower(tabWord[i-j])
					tabWord[i-j] = cap(tabWord[i-j])
				}
				confirm = nbre
			}
			if nbChar != len(tabWord[i+1]) {
				tabWord = append(tabWord[:i], tabWord[i+2:]...)
			}
			if confirm != 0 {
				i--
			}
		} else if tabWord[i] == "(low," && i+1 < len(tabWord) {
			if i > 0 {
				// nbre de Caractère
				for k := 0; k < len(tabWord[i+1]); k++ {
					if tabWord[i+1][k] == '-' || tabWord[i+1][k] == '+' {
						continue
					}
					if !unicode.IsNumber(rune(tabWord[i+1][k])) {
						nbChar++
					}
				}
				// Gestion de la transformation
				nbre, _ := strconv.Atoi(tabWord[i+1][:len(tabWord[i+1])-nbChar])
				if nbre > i {
					nbre = i
				}
				for !isAlphaNum(tabWord[i-y]) {
					if i <= y {
						break
					}
					y++
				}
				for j := y; j <= nbre+(y-1); j++ {
					tabWord[i-j] = strings.ToLower(tabWord[i-j])
				}
				//Surplus
				for l := 1; l < nbChar; l++ {
					surplus += tabWord[i+1][len(tabWord[i+1])-l:]
				}
				tabWord[i-1] += surplus
				confirm = nbre
			}
			if nbChar != len(tabWord[i+1]) {
				tabWord = append(tabWord[:i], tabWord[i+2:]...)
			}
			if confirm != 0 {
				i--
			}
		}
	}
	return tabWord
}