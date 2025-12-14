package main
import (
	"fmt"
	"strings"
)
func PrintGraphicAscii(text string, graphicAscii string) {
	asciiArt := AsciiArt(graphicAscii)
	blocs := [][]string{}
	graphicAsciiChars := []string{}
	graphicAsciichar := ""

	for i := 0; i < len(text); i++ {
		char := text[i]
		if char == '\\' && text[i+1] == 'n' && i+1 < len(text) {
			blocs = append(blocs, graphicAsciiChars)
			graphicAsciiChars = []string{}
			i++
			continue
		}

		for key, value := range asciiArt {
			if char == byte(key) {
				graphicAsciichar = value
				graphicAsciiChars = append(graphicAsciiChars, graphicAsciichar)
				break
			}
		}
	}
	blocs = append(blocs, graphicAsciiChars)
	// Afficher
	for _, b := range blocs {
		if len(b) == 0 {
			fmt.Println()
			continue
		}
		for q := 0; q < 8; q++ {
			for _, r := range b {
				lignes := strings.Split(r, "\n")
				if q < len(lignes) {
					fmt.Print(lignes[q])
				}
			}
			fmt.Println()
		}
	}
}
