package main

func vowel(SplitFile []string) bool {
	vowel := []string{"a", "e", "u", "i", "o", "y"}
	for _, str := range vowel {
		if string(SplitFile[1]) == str {
			return true
		}
	}
	return false
}
