package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	//"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Vous devez entrer un seul Argument  < string > ")
		return
	}

	input := os.Args[1]

	if input == "" {
		return
	}
	if input == "\\n" {
		fmt.Println()
		return
	}

	data, err := http.Get("https://learn.zone01oujda.ma/api/content/root/01-edu_module/content/ascii-art/standard.txt")
	if err != nil {
		fmt.Println("Erreur lors de la Récuperation d API: ", err)
	}

	body, err := io.ReadAll(data.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du flux : ", err)
		return
	}

	// fmt.Println(body)

	content := string(body)

	//fmt.Println(content)
	/*
		API standard => 32-126
		donc je veux  stocker ds une map clé - valeur  ; le cle c'est le ascii courant  et le valeur c'est la représentation graphique
	*/

	// ascii := 32

	// for i := 0; i+7 < len(content); i += 7 {
	// 	fmt.Println(content[i:i+8])
	// }

	//asciiArt := map[int]string{}
	lignes := []string{}
	ascii := 32
	jump := 0
	for _, r := range content {
		if jump == 8 {
			ascii++
		}	
		if r == '\n' {
			jump++
		}
	}

	fmt.Println(strings.Join(lignes, "stop"))
}
