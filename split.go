package main

import "strings"

func SplitText(file string) []string {
	return strings.Split(string(file), " ")
}