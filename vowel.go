package main

func vowel(SplitFile []string) []string {
	vowels := "aeuiohAEIOUH"
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
	return SplitFile
}
