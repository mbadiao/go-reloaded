package main

import "regexp"

func Par(input string) string {
	// Créer une expression régulière pour rechercher '(' ou ')'
	//enlever les espace a l'interieur des parentheses
	re := regexp.MustCompile(`\(\s`)
	re1 := regexp.MustCompile(`\s\)`)
	// re2 := regexp.MustCompile(`\s\(`)
	// re3 := regexp.MustCompile(`\)\s`)
	//si on a pas d'espace entre ')(' => ') ('
	re4 := regexp.MustCompile(`\)\(`)
	//si on a pas d'espace apres la parenthese fermante ')' => ') '
	re5 := regexp.MustCompile(`\)`)
	//si on a pas d'espace apres la parenthese ouvrante '(' => ' ('
	re6 := regexp.MustCompile(`\(`)
	// Remplacer chaque parenthèse '(' ou ')' par un espace avant et après
	result := re.ReplaceAllString(input, "(")
	result = re1.ReplaceAllString(input, ")" )
	// result = re2.ReplaceAllString(input, " (")
	// result = re3.ReplaceAllString(input, ") ")
	result = re4.ReplaceAllString(input, ") (")
	result = re5.ReplaceAllString(input, ") ")
	result = re6.ReplaceAllString(input, " (")

	return result
}

//func repar(input)