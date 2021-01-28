package functions

//var t *template.Template

func PrintAscii(eachTextSplit string, asciiLines []string) string {

	//var w http.ResponseWriter

	characterLine := 0

	var finalStr = ``

	for everyLines := 0; everyLines < 9; everyLines++ { // Boucle permettant de mettre chaque ligne

		str := ""

		for _, char := range eachTextSplit { //Boucle regardant chaque string du tableau de string

			characterLine = (int(char) - 32) * 9             //cherche la première ligne du caractère correspondant
			str = str + asciiLines[characterLine+everyLines] // concaténation de la ligne des différents caracères

		}
		// fmt.Println(str)

		finalStr = finalStr + str + `<br>`
	}

	return finalStr
}
