package main

import (
	"fmt"
	"os"
	"strings"
)

func OpenFiles(TextFile string) []string {
	file, err := os.ReadFile(TextFile)
	if err != nil {
		fmt.Println("erreur l'ors de l'ouverture du fichier", err)
	}
	if !TextGood(string(TextFile)) {
		return []string{ "Your", "file content a irregular word" }
	}
	if len(string(file)) == 0 {
		return []string{ "Your","file is empty !!!" }
	}

	fileFinal := strings.Fields(string(file))
	return fileFinal
}
