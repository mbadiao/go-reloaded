package main

import (
	"fmt"
	"os"
//	"strconv"
//	"strings"
)

func main() {
	arg := os.Args[1:] 
// if they are more than 2 arguments
	if len(arg) != 2 {
		os.Exit(1)
	} else {
		fmt.Println("very")
	}
}

