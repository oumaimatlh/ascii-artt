package main

func AsciiArt(graphicAscii string) map[int]string {
	asciiArt := map[int]string{}
	ascii := 32
	var graphic string
	jump := 0
	check := false

	for i := 0; i < len(graphicAscii); i++ {
		if graphicAscii[i] == '\r' && i+1 < len(graphicAscii) && graphicAscii[i+1] == '\n' {
			i++ 
		}
		if graphicAscii[i] != '\n' {
			graphic += string(graphicAscii[i])
			check = true
		} else if check {
			if jump == 8 {
				asciiArt[ascii] = graphic
				ascii++
				jump = 0
				graphic = ""
			} else {
				graphic += string(graphicAscii[i])
				jump++
			}
		}
	}
	if graphic != "" && jump == 8 {
		asciiArt[ascii] = graphic
	}
	return asciiArt
}