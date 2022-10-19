package basics

import "os"

func FileVar(file string) string {
	textraw, err := os.ReadFile(file)
	if err != nil {
		os.Exit(0)
	}
	return string(textraw)
}
