package main

import (
	"fmt"
	"os"
 	"strings"
	"bufio"
)

func main() {
	arg := os.Args[1:] 
// if they are more than 2 arguments
	if len(arg) > 2 {
		fmt.Println("Il ne doit  pas y avoir plus de deux arguments")
		os.Exit(1)
	}
	if strings.Contains(arg[0],".txt") {
		 file,err := os.Open(arg[0])
		 if err != nil {
			fmt.Println("erreur l'ors de l'ouverture du fichier",err)
		 }
		 defer file.Close() // Assurez-vous de fermer le fichier à la fin de l'exécution de la fonction

		 // Créer un scanner pour lire le fichier ligne par ligne
		 scanner := bufio.NewScanner(file)
	 
		 // Parcourir le fichier ligne par ligne
		 for scanner.Scan() {
			 // La ligne actuelle se trouve dans scanner.Text()
			 line := scanner.Text()
			fmt.Println(line)
		 }
	}
}
