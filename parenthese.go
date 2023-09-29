package main

import "regexp"

func Par(input string) []string {
	// Créer une expression régulière pour rechercher '(' ou ')'
	//enlever les espace a l'interieur des parentheses
	re3 := regexp.MustCompile(`\(\s`)
	result := re3.ReplaceAllString(input, "(")
	re1 := regexp.MustCompile(`\s\)`)
	result = re1.ReplaceAllString(input, ")" )
	// //si on a pas d'espace entre ')(' => ') ('
	re4 := regexp.MustCompile(`\)\(`)
	re7 := regexp.MustCompile(`\'\(`)
	re8 := regexp.MustCompile(`\)\'`)
	result = re7.ReplaceAllString(input, "' (")
	result = re8.ReplaceAllString(input, ") '")
	result = re4.ReplaceAllString(input, ") (")
	re5 := regexp.MustCompile(`\(\(`)
	result = re5.ReplaceAllString(input, "( (")
	re6 := regexp.MustCompile(`\)\)`)
	result = re6.ReplaceAllString(input, ") )")
	// //si on a pas d'espace apres la parenthese fermante ')' => ') '
	//re := regexp.MustCompile(`(\S+)([()])|([()])(\S+)`)

	// Remplacer chaque occurrence de "(" ou ")" collée par "$1 $2" où "$1" est le mot précédent et "$2" est la parenthèse.
//	result := re.ReplaceAllString(input, "$1 $2$3 $4")
	// Remplacer chaque parenthèse '(' ou ')' par un espace avant et après
	
	
	
	 
	
	//re := regexp.MustCompile(`(\S+)\(`)
	//re1 := regexp.MustCompile(`\)(\S+)`)
	//re1 := regexp.MustCompile(`((\S+)\`)
	//result := re.ReplaceAllString(input, ") $1")
	//result := re.ReplaceAllString(input, "$1 (")
	//result = re1.ReplaceAllString(input, ") $1")
	return SplitText(result)
}

func repar(input string) string {
	re7 := regexp.MustCompile(`\s\(`)
	result := re7.ReplaceAllString(input, "(")

	return result
}