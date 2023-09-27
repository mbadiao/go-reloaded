package main

import "regexp"

func Par(input string) string {
	// Créer une expression régulière pour rechercher '(' ou ')'
	//enlever les espace a l'interieur des parentheses
	 //re3 := regexp.MustCompile(`\(\s`)
	// re1 := regexp.MustCompile(`\s\)`)
	
	// //si on a pas d'espace entre ')(' => ') ('
	// re4 := regexp.MustCompile(`\)\(`)
	// //si on a pas d'espace apres la parenthese fermante ')' => ') '
	//re := regexp.MustCompile(`(\S+)([()])|([()])(\S+)`)

	// Remplacer chaque occurrence de "(" ou ")" collée par "$1 $2" où "$1" est le mot précédent et "$2" est la parenthèse.
//	result := re.ReplaceAllString(input, "$1 $2$3 $4")
	// Remplacer chaque parenthèse '(' ou ')' par un espace avant et après
	//result = re3.ReplaceAllString(input, "(")
	//result = re1.ReplaceAllString(input, ")" )
	// result = re4.ReplaceAllString(input, ") (")
	
	re := regexp.MustCompile(`(\S+)\(`)
	re1 := regexp.MustCompile(`\)(\S+)`)
	//re1 := regexp.MustCompile(`((\S+)\`)
	//result := re.ReplaceAllString(input, ") $1")
	result := re.ReplaceAllString(input, "$1 (")
	result = re1.ReplaceAllString(input, ") $1")
	return result
}

func repar(input string) string {
	re7 := regexp.MustCompile(`\s\(`)
	result := re7.ReplaceAllString(input, "(")

	return result
}