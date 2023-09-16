package main

func vowel(SplitFile []string) []string {
	vowel := "aeuiohAEIOUH"
	for i := 0; i < len(SplitFile); i++ {
		if SplitFile[i] == "a" || SplitFile[i] == "A" {
			if i < len(SplitFile)-1 {
				for j := 0; j < len(SplitFile); j++ {
					if SplitFile[i+1][0] == vowel[j] {
						SplitFile[i] += "n"
					}
				}
			}
		}
	}
	return SplitFile
}