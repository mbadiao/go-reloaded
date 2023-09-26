package main

import (
	"regexp"
	"strings"
)

func aps(s string) string {
	preResult := regexp.MustCompile(`' (.*?) '`)
	result := preResult.ReplaceAllString(s, "'$1'")
	text := strings.Fields(result)
	return strings.Join(text, " ")
}
