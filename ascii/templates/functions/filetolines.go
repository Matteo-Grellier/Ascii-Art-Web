package functions

import (
	"bufio"
	"os"
)

//lecture du fichier ascii.txt dans lequel se trouve tous les caractères en ASCII-Art
//Puis mise de chaque ligne dans un tableau de string

func FilesToLines(filePath string) (asciiArtLines []string) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close() //fermer le fichier au moment du "return"

	scanner := bufio.NewScanner(file) //
	for scanner.Scan() {
		asciiArtLines = append(asciiArtLines, scanner.Text())
	}
	//attention ici il faut vérifier si c'est bon (car la c'est pas fait)
	err = scanner.Err()
	if err != nil {
		return nil
	}
	return asciiArtLines

}
