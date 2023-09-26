package main

func verify(s string) bool {
	var nbr int
	for _, char := range s {
		if char < 48 || char > 57 && char < 65 || char > 90 && char < 97 || char > 122 {
			nbr++
		}
	}
	return nbr != len(s)
}