package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	// if they are more than 2 arguments
	if len(arg) == 2 {
		fmt.Println("Il ne doit  pas y avoir plus de deux arguments")
		os.Exit(1)
	}
	fmt.Println(OpenFiles(arg[0]))
}
