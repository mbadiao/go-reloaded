package main

import "strings"

func vowel(File string) string {
	vowels := "aeuiohAEIOUH"
	SplitFile := strings.Fields( File)
	for i := 0; i < len(SplitFile); i++ {
		if SplitFile[i] == "a" || SplitFile[i] == "A" {
			if i < len(SplitFile)-1 {
				for j := 0; j < len(vowels); j++ {
					if len(SplitFile[i+1]) > 0 && SplitFile[i+1][0] == vowels[j] {
						SplitFile[i] += "n"
						break // Exit the inner loop once "n" is added
					}
				}
			}
		}
	}
	return strings.Join(SplitFile," ")
}
