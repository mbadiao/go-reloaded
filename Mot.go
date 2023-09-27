package main

import (
	"strings"
	"unicode"
)
// Distinction des mots clés
func GoodKeys(tabWord []string) []string {
	var nbNum int
	for i := 0; i < len(tabWord); i++ {
		if strings.Contains(tabWord[i], "(up)") {
			for j := 1; j < len(tabWord[i])-1; j++ {
				if tabWord[i][j] == '(' && tabWord[i][j-1] != ' ' {
					tabWord[i] = tabWord[i][:j] + " " + tabWord[i][j:]
				} else if tabWord[i][j] == ')' && tabWord[i][j+1] != ' ' {
					tabWord[i] = tabWord[i][:j+1] + " " + tabWord[i][j+1:]
				}
			}
		}
		if strings.Contains(tabWord[i], "(low)") {
			for j := 1; j < len(tabWord[i])-1; j++ {
				if tabWord[i][j] == '(' && tabWord[i][j-1] != ' ' {
					tabWord[i] = tabWord[i][:j] + " " + tabWord[i][j:]
				} else if tabWord[i][j] == ')' && tabWord[i][j+1] != ' ' {
					tabWord[i] = tabWord[i][:j+1] + " " + tabWord[i][j+1:]
				}
			}
		}
		if strings.Contains(tabWord[i], "(cap)") {
			for j := 1; j < len(tabWord[i])-1; j++ {
				if tabWord[i][j] == '(' && tabWord[i][j-1] != ' ' {
					tabWord[i] = tabWord[i][:j] + " " + tabWord[i][j:]
				} else if tabWord[i][j] == ')' && tabWord[i][j+1] != ' ' {
					tabWord[i] = tabWord[i][:j+1] + " " + tabWord[i][j+1:]
				}
			}
		}
		if strings.Contains(tabWord[i], "(bin)") {
			for j := 1; j < len(tabWord[i])-1; j++ {
				if tabWord[i][j] == '(' && tabWord[i][j-1] != ' ' {
					tabWord[i] = tabWord[i][:j] + " " + tabWord[i][j:]
				} else if tabWord[i][j] == ')' && tabWord[i][j+1] != ' ' {
					tabWord[i] = tabWord[i][:j+1] + " " + tabWord[i][j+1:]
				}
			}
		}
		if strings.Contains(tabWord[i], "(hex)") {
			for j := 1; j < len(tabWord[i])-4; j++ {
				if tabWord[i][j:j+5] == "(hex)" && tabWord[i][j-1] != ' ' {
					tabWord[i] = tabWord[i][:j] + " " + tabWord[i][j:]
					break
				}
			}
			for j := 1; j < len(tabWord[i])-1; j++ {
				if tabWord[i][j] == ')' && tabWord[i][j+1] != ' ' {
					tabWord[i] = tabWord[i][:j+1] + " " + tabWord[i][j+1:]
					break
				}
			}
		}
		if strings.Contains(tabWord[i], "(up,") {
			nbNum = 0
			if i+1 < len(tabWord) {
				for k := 0; k < len(tabWord[i+1]); k++ {
					if tabWord[i+1][k] == '-' || tabWord[i+1][k] == '+' {
						continue
					}
					if unicode.IsNumber(rune(tabWord[i+1][k])) {
						nbNum++
						break
					}
				}
			}
			if nbNum > 0 {
				for j := 1; j < len(tabWord[i])-3; j++ {
					if tabWord[i][j:j+4] == "(up," && tabWord[i][j-1] != ' ' {
						tabWord[i] = tabWord[i][:j] + " " + tabWord[i][j:]
						break
					} /* else if tabWord[i+1][j] == ')' && tabWord[i+1][j+1] != ' ' {
						tabWord[i+1] = tabWord[i+1][:j+1] + " " + tabWord[i+1][j+1:]
					} */
				}
				if i+1 < len(tabWord) {
					for j := 1; j < len(tabWord[i+1])-1; j++ {
						/* if tabWord[i][j] == '(' && tabWord[i][j-1] != ' ' {
							tabWord[i] = tabWord[i][:j] + " " + tabWord[i][j:]
						} */if tabWord[i+1][j] == ')' && tabWord[i+1][j+1] != ' ' {
							tabWord[i+1] = tabWord[i+1][:j+1] + " " + tabWord[i+1][j+1:]
							break
						}
					}
				}
			}
		}
		if strings.Contains(tabWord[i], "(cap,") {
			nbNum = 0
			if i+1 < len(tabWord) {
				for k := 0; k < len(tabWord[i+1]); k++ {
					if tabWord[i+1][k] == '-' || tabWord[i+1][k] == '+' {
						continue
					}
					if unicode.IsNumber(rune(tabWord[i+1][k])) {
						nbNum++
						break
					}
				}
			}
			if nbNum > 0 {
				for j := 1; j < len(tabWord[i])-4; j++ {
					if tabWord[i][j:j+5] == "(cap," && tabWord[i][j-1] != ' ' {
						tabWord[i] = tabWord[i][:j] + " " + tabWord[i][j:]
						break
					}
				}
				if i+1 < len(tabWord) {
					for j := 1; j < len(tabWord[i+1])-1; j++ {
						if tabWord[i+1][j] == ')' && tabWord[i+1][j+1] != ' ' {
							tabWord[i+1] = tabWord[i+1][:j+1] + " " + tabWord[i+1][j+1:]
							break
						}
					}
				}
			}
		}
		if strings.Contains(tabWord[i], "(low,") {
			nbNum = 0
			if i+1 < len(tabWord) {
				for k := 0; k < len(tabWord[i+1]); k++ {
					if tabWord[i+1][k] == '-' || tabWord[i+1][k] == '+' {
						continue
					}
					if unicode.IsNumber(rune(tabWord[i+1][k])) {
						nbNum++
						break
					}
				}
			}
			if nbNum > 0 {
				for j := 1; j < len(tabWord[i])-4; j++ {
					if tabWord[i][j:j+5] == "(low," && tabWord[i][j-1] != ' ' {
						tabWord[i] = tabWord[i][:j] + " " + tabWord[i][j:]
						break
					}
				}
				if i+1 < len(tabWord) {
					for j := 1; j < len(tabWord[i+1])-1; j++ {
						if tabWord[i+1][j] == ')' && tabWord[i+1][j+1] != ' ' {
							tabWord[i+1] = tabWord[i+1][:j+1] + " " + tabWord[i+1][j+1:]
							break
						}
					}
				}
			}
		}
	}
	tabWord = strings.Fields(strings.Join(tabWord, " "))
	return tabWord
}