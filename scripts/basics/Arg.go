package basics

import (
	"fmt"
	"os"
)

func PrintUsage() { // impression du message d'usage
	fmt.Println("Usage: go run . [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . something standard")
	os.Exit(0)
}

func Arg(nARGS, expargs int) (input string, banner string, output string) { // vérification d'erreurs sur les arguments
	if nARGS > expargs { // nombre d'arguments supérieur à 2
		PrintUsage()
	}
	if nARGS < 1 {
		PrintUsage()
	}

	input = string(os.Args[1]) // caractères enre 32 et 126
	for _, char := range input {
		if char < 32 || char > 126 {
			PrintUsage()
		}
	}
	if nARGS == 2 {
		if os.Args[2] != "standard" && os.Args[2] != "shadow" && os.Args[2] != "thinkertoy" { // format différent de "standard", "shadow" et "thinkertoy"
			PrintUsage()
		}
		banner = "./banner/" + string(os.Args[2]) + ".txt"
	} else {
		banner = "./banner/standard.txt"
	}
	// if banner == "./banner.txt"
	// output = string(os.Args[3])
	// splitter := strings.Split(output, "=")
	// output = string(splitter[1])
	return input, banner, output
}
