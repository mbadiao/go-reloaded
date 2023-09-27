package main

import (
	"regexp"
//	"strings"
)

func aps(s string) string {
	preResult := regexp.MustCompile(`'\s*(.*?)\s*'`)
	result := preResult.ReplaceAllString(s, "'$1'")
	//text := strings.Fields(result)
	//return strings.Join(text, " ")
	return result
}
