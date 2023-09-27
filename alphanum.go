package main

func isAlphaNum(s string) bool {
	var count int
	for _, char := range s {
		if char < 48 || char > 57 && char < 65 || char > 90 && char < 97 || char > 122 {
			count++
		}
	}
	return count != len(s)
}