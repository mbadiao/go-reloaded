package main

import (
	"fmt"
	"os"
)

func main() {
	// if they are more than 2 arguments
	if len(os.Args[1:]) > 2 {
		fmt.Println("Il ne doit  pas y avoir plus de deux arguments")
		os.Exit(1)
	}
	x := OpenFiles(os.Args[1])
	y := vowel(SwitchFunc(SplitText(x)))
	fmt.Println(y)
}
