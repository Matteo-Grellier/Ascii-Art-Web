package main

import (
	"fmt"
	"os"
	"strings"

	asciifunc ".."
)

func main() {

	if len(os.Args)-1 == 2 {

		textToConvert := os.Args[1]
		var asciiLines []string

		if os.Args[2] == "standard" || os.Args[2] == "shadow" || os.Args[2] == "thinkertoy" {

			nameOfFiles := os.Args[2] + ".txt"
			asciiLines = asciifunc.FilesToLines("../" + nameOfFiles) // Lecture / récupération des lignes du fichier avec les caractères en ascii-art
		} else {
			asciiLines = asciifunc.FilesToLines("") // dans le cas où l'utilisateur entre une mauvais nom de fichier.
		}

		arrayOfSplitStrings := strings.Split(textToConvert, "\\n") //Tableau de string séparant les mots s'il y a des '\n'

		for _, eachTextSplit := range arrayOfSplitStrings { //Boucle parcourant le tableau de string

			//Condition de vérification s'il y a eu une erreur (si c'est faux, on continue le programme)...
			if asciiLines == nil {
				fmt.Println("Error reading ASCII file...")
				break
			} else {
				asciifunc.PrintAscii(eachTextSplit, asciiLines)
			}
		}

	} else {

		fmt.Println("Sorry, but there is not the good number of argument")
	}

}
