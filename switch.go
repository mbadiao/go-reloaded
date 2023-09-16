package main

import (
	"strconv"
)

func SwitchFunc(SplitFile []string) []string {
	for i := 0; i < len(SplitFile); i++ {
		switch SplitFile[i] {
		case "(hex)":
			if i > 0 {
				SplitFile[i-1] = hex(SplitFile[i-1])
			}
			// Supprime le modèle actuel et l'élément suivant
			SplitFile = append(SplitFile[:i], SplitFile[i+2:]...)
			i-- // Décrémenter i pour rester à la même position après la suppression
		case "(bin)":
			if i > 0 {
				SplitFile[i-1] = bin(SplitFile[i-1])
			}
			SplitFile = append(SplitFile[:i], SplitFile[i+1:]...)
			i--
		case "(up)":
			if i > 0 {
				SplitFile[i-1] = up(SplitFile[i-1])
			}
			SplitFile = append(SplitFile[:i], SplitFile[i+1:]...)
			i--
		case "(low)":
			if i > 0 {
				SplitFile[i-1] = low(SplitFile[i-1])
			}
			SplitFile = append(SplitFile[:i], SplitFile[i+1:]...)
			i--
		case "(cap)":
			if i > 0 {
				SplitFile[i-1] = cap(SplitFile[i-1])
			}
			SplitFile = append(SplitFile[:i], SplitFile[i+2:]...)
			i--
		case "(low,":
			nbr, _ := strconv.Atoi(SplitFile[i+1][:len(SplitFile[i+1])-1])
			if i > 0 {
				for j := 1; j <= nbr; j++ {
					SplitFile[i-j] = low(SplitFile[i-j])
				}
			}
			SplitFile = append(SplitFile[:i], SplitFile[i+2:]...)
			i--
		case "(cap,":
			nbr, _ := strconv.Atoi(SplitFile[i+1][:len(SplitFile[i+1])-1])
			if i > 0 {
				for j := 1; j <= nbr; j++ {
					SplitFile[i-j] = cap(SplitFile[i-j])
				}
			}
			SplitFile = append(SplitFile[:i], SplitFile[i+2:]...)
			i--
		case "(up,":
			nbr, _ := strconv.Atoi(SplitFile[i+1][:len(SplitFile[i+1])-1])
			if i > 0 {
				for j := 1; j <= nbr; j++ {
					SplitFile[i-j] = up(SplitFile[i-j])
				}
			}
			SplitFile = append(SplitFile[:i], SplitFile[i+2:]...)
			i--
		}
	}
	return SplitFile
}
