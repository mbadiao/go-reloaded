package main

func TextGood(s string) bool {
	for _, char := range s {
		if char < 0 || char > 127 {
			return false
		}
	}
	return true
}