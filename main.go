package main

import (
	"fmt"
	"os"
	"strings"
	//"strings"
)

func main() {
	// if they are more than 2 arguments
	if len(os.Args[1:]) != 2 {
		fmt.Println("Il ne doit y avoir plus de deux arguments")
		os.Exit(1)
	}

	TextFile := OpenFiles(os.Args[1])
	AfterPonctuations := punc(TextFile)
//	fmt.Println(AfterPonctuations)
	stringsF := strings.Join(AfterPonctuations,  " ")
	AfterPonctuations = Par(stringsF)
	SplitAfterPonctuations := GoodKeys(AfterPonctuations)
	//fmt.Println(SplitAfterPonctuations)
	AfterPonctuations = punc(SplitAfterPonctuations)
	ApsText := aps(AfterPonctuations)
	AfterPonctuations = Par(stringsF)
	AfterKey := GoodKeys(ApsText)
	ModifText := SwitchFunc(AfterKey)
//	fmt.Println(ModifText)
	AfterModifText := hexbin(ModifText)
	AfterModifText = punc(AfterModifText)
	AfterModifText = aps(AfterModifText)
//	fmt.Println(AfterModifText)
	AfterModifText = vowel(AfterModifText)
	AfterModifText = GoodKeys(AfterModifText)
	AfterModifText = SwitchFunc(AfterModifText)
	AfterModifText = punc(AfterModifText)
	os.WriteFile(os.Args[2], []byte(strings.Join(AfterModifText, " ")), 0666)
	
}
