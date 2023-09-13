package main

import (
	"fmt"
	"os"
	"strings"
)

func OpenFiles(s string) []string {
	file, err := os.ReadFile(s)
	if err != nil {
		fmt.Println("erreur l'ors de l'ouverture du fichier", err)
	}
	FileSplit := strings.Split(string(file)," ")
	return FileSplit
}
