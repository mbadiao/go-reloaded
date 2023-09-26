package main

import "regexp"

func punc(file string) string {
	//pour supprimer l'espace avant la ponctuation
	re := regexp.MustCompile(`\s+([.,;:!?]+)`)
	
	final := re.ReplaceAllString(file, "$1")
	
	return final
}

func repunc(file string) string{
	re1 := regexp.MustCompile(`([.,;:!?]+)`)
	final := re1.ReplaceAllString(file, "$1 ")
	return final
}