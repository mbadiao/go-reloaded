package main

import (
	"fmt"
	"os"
)

func OpenFiles(s string) string {
	file, err := os.ReadFile(s)
	if err != nil {
		fmt.Println("erreur l'ors de l'ouverture du fichier", err)
	}
	if len(string(file)) == 0 {
		return "your file is empty !!!"
	}
	return string(file)
}
