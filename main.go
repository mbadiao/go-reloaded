package main

import (
	"fmt"
	"os"
)

func main() {
	// if they are more than 2 arguments
	if len(os.Args[1:]) != 1 {
		fmt.Println("Il ne doit y avoir deux arguments")
		os.Exit(1)
	}

	x := OpenFiles(os.Args[1])
	x = punc(x)
	y := vowel(SwitchFunc(x))
	punctfile := punc(y)
	fmt.Println(punctfile)
}
