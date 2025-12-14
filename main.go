package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Vous devez entrer un seul Argument  < string > ")
		return
	}
	inputText := os.Args[1]
	if inputText == "" {
		return
	}
	if inputText == "\\n" {
		fmt.Println()
		return
	}
	for _, r := range inputText {
		if !(r >= 32 && r <= 126) {
			fmt.Println("Input Non validée : les caractéres doivent etre en Ascii 32-126")
			return
		}
	}
	data, err := os.ReadFile("./bannersFiles/standard.txt")
	if err != nil {
		fmt.Println("Ereur au niveau du fichier: ", err)
		return
	}
	graphicAscii := string(data)
	PrintGraphicAscii(inputText, graphicAscii)
}
