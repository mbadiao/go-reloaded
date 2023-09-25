package main

import "strings"

func rm(s string) string {
	word := strings.Fields(s)
	result := strings.Join(word, " ") 
	return result
}