package ascii

import (
	// "fmt"
	"strings"
)

func ConvertAsciiArt(splitterAAT []string, input string) (asciiart string) {
	splitterInput := strings.Split(input, "\\n") // Splitter chaque chaîne de caractère lorsqu'il y a un "\n"
	//													Que faire si on veut afficher un \n ?
	for _, eachLine := range splitterInput { // 	Parcourir chaque ligne de splitterInput
		var arrascii [][]string

		for _, char := range eachLine {

			ncharline := (char-32)*9 + 1                       // N° de ligne du caractère
			asciiLines := splitterAAT[ncharline : ncharline+8] // prendre les 8 lignes
			arrascii = append(arrascii, asciiLines)            // les stocker dans une ligne du tableau
			// fmt.Println(asciiLines)
		}
		for iline := 0; iline < 8; iline++ { // 	boucler 8x
			for _, line := range arrascii { // 			Parcourir les tableaux des caractères
				asciiart += line[iline] // 					Ajouter la n-ième ligne d'un caractère
			}
			asciiart += "\n" // 						retour à la ligne après avoir ajouté la n-ième ligne du dernier caractère à la fin de la ligne
		}

	}
	return asciiart
}

/*


   Gestion du saut de ligne
   if i >= 2 && nchar == 'n' && inputi[i-1] == '\\' && inputi[i-2] != "\\" {
   	arrascii = append(arrascii, "\n")
   }
  stocker chacun des caractères dans un tableau de tableaux (un caractère par ligne et chaque ligne d'un caractère dans une colonne) :
[ [L1] [L2] [L3] [L4] [L5] [L6] [L7] [L8]  --> Caractère A
  [L1] [L2] [L3] [L4] [L5] [L6] [L7] [L8]  --> caractère \
  [L1] [L2] [L3] [L4] [L5] [L6] [L7] [L8]] --> caractère n
*/
