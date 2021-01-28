package functions

// PrintAscii permet l'écriture (sur une seule ligne) de la chaine de caractère entrée par l'utilisateur.
func PrintAscii(eachTextSplit string, asciiLines []string) string {

	characterLine := 0

	var finalStr = ``

	for everyLines := 0; everyLines < 9; everyLines++ { // Boucle permettant de mettre chaque ligne

		str := ""

		for _, char := range eachTextSplit { //Boucle regardant chaque string du tableau de string

			characterLine = (int(char) - 32) * 9             //cherche la première ligne du caractère correspondant
			str = str + asciiLines[characterLine+everyLines] // concaténation de la ligne des différents caractères

		}

		finalStr = finalStr + str + `<br>` //Concaténation dans une seule string (retour à la ligne possible grâce à la balise "<br>"
	}

	return finalStr
}
