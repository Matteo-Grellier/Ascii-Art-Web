package functions

import (
	"fmt"
	"strings"
	// asciifunc "../functions"
)

func Ascii(textToConvert string, banner string) string {

	var asciiLines []string
	var finalStr string
	var buffStr string

	if banner == "standard" || banner == "shadow" || banner == "thinkertoy" {

		nameOfFiles := banner + ".txt"
		asciiLines = FilesToLines("../templates/" + nameOfFiles) // Lecture / récupération des lignes du fichier avec les caractères en ascii-art
	} else {
		asciiLines = FilesToLines("") // dans le cas où l'utilisateur entre une mauvais nom de fichier.
	}

	arrayOfSplitStrings := strings.Split(textToConvert, "\\n") //Tableau de string séparant les mots s'il y a des '\n'

	for _, eachTextSplit := range arrayOfSplitStrings { //Boucle parcourant le tableau de string

		//Condition de vérification s'il y a eu une erreur (si c'est faux, on continue le programme)...
		if asciiLines == nil {
			fmt.Println("Error reading ASCII file...")
			break
		} else {
			buffStr = PrintAscii(eachTextSplit, asciiLines)
		}

		finalStr = finalStr + buffStr
	}

	return finalStr
}
