package main

func cap(s string) string {
	all_rune := []rune(s)
	compt := 1
	for i := range all_rune {
		if (all_rune[i] <= 'z' && all_rune[i] >= 'a') || (all_rune[i] <= 'Z' && all_rune[i] >= 'A') || (all_rune[i] <= '9' && all_rune[i] >= '0') {
			if compt == 1 && all_rune[i] <= 'z' && all_rune[i] >= 'a' {
				all_rune[i] = all_rune[i] - 32
			}
			if compt > 1 && all_rune[i] <= 'Z' && all_rune[i] >= 'A' {
				all_rune[i] = all_rune[i] + 32
			}
			compt++
		}
	}
	return string(all_rune)
}
