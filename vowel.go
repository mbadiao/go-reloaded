package main

import "unicode"

func vowel(tabWord []string) []string {
	str := "aeiouhAEIOUH"
	var k int
	for i := 0; i < len(tabWord); i++ {
		if tabWord[i] == "a" || tabWord[i] == "A" {
			if i < len(tabWord)-1 {
				for j := 0; j < len(str); j++ {
					if !unicode.IsLetter(rune(tabWord[i+1][k])) {
						if k < len(tabWord[i+1]) {
							k++
						}
						if k == len(tabWord[i+1]) {
							break
						}
						if j > 0 {
							j--
						}
					}
					if tabWord[i+1][k] == str[j] {
						tabWord[i] += "n"
						break
					}
				}
			}
		}
	}
	return tabWord
}
