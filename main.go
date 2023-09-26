package main

import (
	"fmt"
	"os"
	"strings"
	//"strings"
)

func main() {
	// if they are more than 2 arguments
	if len(os.Args[1:]) != 1 {
		fmt.Println("Il ne doit y avoir deux arguments")
		os.Exit(1)
	}

	x := OpenFiles(os.Args[1])
	x = punc(x)
	x = aps(x)
	sx := strings.Split(x, " ")
	modif := SwitchFunc(sx)
	y := vowel(modif)
	punctfile := punc(y)
	repunctfile := repunc(punctfile)
	resx := strings.Split(repunctfile," ")
	remodif := SwitchFunc(resx)
	remodif = aps(remodif)
	fmt.Println(remodif)
}
