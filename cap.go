package main

import "strings"

func cap(s string) string {
	all_rune := []rune(s)
	compt := 0
	for i := range all_rune {
		if (all_rune[i] <= 'z' && all_rune[i] >= 'a') || (all_rune[i] <= 'Z' && all_rune[i] >= 'A') || (all_rune[i] <= '9' && all_rune[i] >= '0') {
			compt++
		} else {
			compt = 0
		}
		if compt == 1 && all_rune[i] <= 'z' && all_rune[i] >= 'a' {
			all_rune[i] = all_rune[i] - 32
		}
		if compt > 1 && all_rune[i] <= 'Z' && all_rune[i] >= 'A' {
			all_rune[i] = all_rune[i] + 32
		}
	}
	var all_rune1 string
	for i := 0; i < len(all_rune); i++ {
		if all_rune[i] == 39 {
			all_rune1 = string(all_rune[:i]) + (strings.ToLower(string(all_rune[i:])))
		}
	}
	return all_rune1
}
