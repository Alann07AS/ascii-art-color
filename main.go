package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	color := map[string]string{ // Creating Key/Value color
		"reset":  "\033[0m",
		"black":  "\033[30m",
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"purple": "\033[35m",
		"cyan":   "\033[36m",
		"white":  "\033[37m",
		"orange": "\033[38;5;208m",
	}

	const separator = "|"
	arg := os.Args[1:]
	if len(arg) != 2 {
		fmt.Println("Usage : go run . [STRING] [OPTION]")
		fmt.Println("EX: go run . something --color=<color>")
		return
	}

	str := arg[0]
	for _, val := range str {
		if (val < 32) || val > 127 {
			fmt.Println("Les Caractères accentués ne sont pas permis")
			return
		}
	}

	empty := false
	for _, val := range str {
		if val != '|' && val != ' ' {
			empty = true
			break
		}
	}
	if !empty {
		fmt.Println("Aucun caractère saisi")
		return
	}

	if strings.Count(str, "|") > 2 {
		fmt.Println("Délimiteur '|' > 2")
		return
	}

	const option = "--color="
	choosenColor := color["yellow"] // default color if not defined
	if arg[1][0:len(option)] == option {
		if choosen := strings.ToLower(arg[1][len(option):]); choosen != "" {
			if _, ok := color[choosen]; !ok { // if color not exist in the map
				fmt.Println("Couleur non valide")
				return
			} else {
				choosenColor = color[choosen]
			}
		}
	} else {
		fmt.Println("Syntaxe incorrecte pour la couleur : --color=")
		return
	}

	file, err := ioutil.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
	}
	fontStr := strings.Split(string(file), "\n")

	// fmt.Println(choosenColor)

	first := strings.Index(str, "|")
	last := strings.LastIndex(str, "|")
	if (first == -1 && last == -1) || first == last { // one or separator not indicated
		last = len(str)
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < len(str); j++ {
			if string(str[j]) == separator {
				continue
			}
			if j <= first || j >= last {
				fmt.Print(color["white"], fontStr[int(str[j]-32)*9+1+i])
			} else {
				fmt.Print(choosenColor, fontStr[int(str[j]-32)*9+1+i])
			}
		}
		fmt.Println()
	}
	fmt.Print(color["reset"]) // if omitted the color of line in the terminal is the last color
}
