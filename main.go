package main

import (
	"fmt"
	"os"
	"strings"

	"asciiart/scripts/ascii"
	"asciiart/scripts/basics"
)

func main() {
	input, banner, _ := basics.Arg(len(os.Args)-1, 2) // stockage et conversion de l'argument en int
	// fmt.Println(input, banner, output)

	switch input { // Traitement des cas où l'argument est un saut de ligne ou un argument vide
	case "\\n":
		fmt.Println()
		os.Exit(0)

	case "":
		os.Exit(0)
	}

	asciiArtTable := basics.FileVar(banner) // Stocker le contenu du fichier dans une variable
	splitterAAT := strings.Split(asciiArtTable, "\n")
	for isAAT, line_sAAT := range splitterAAT {
		fmt.Println(isAAT, line_sAAT)
	}
	asciiart := ascii.ConvertAsciiArt(splitterAAT, input)
	fmt.Print(asciiart)

	/*
	   Argument "\n" ou "" OK

	   range argument
	   	convertir la rune en byte
	   	si la rune est incluse dans le fichier ([32:126]):
	   		nchar = (valeur de la rune en byte - 32) x9 +1
	   		range du splitterAAT
	   			append asciiline de splitterAAT
	   			si ligne vide : break
	   		range  asciiLine
	   			append arrascii de asciiLine


	*/
}
