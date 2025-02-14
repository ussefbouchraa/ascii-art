package asciiartcolor

import (
	"bufio"
	"os"
	S "strings"
)

func InsertValue(scanner *bufio.Scanner) [8]string {
	ArtValue := [8]string{}

	for cp := 0; cp < 8 && scanner.Scan(); cp++ {
		ArtValue[cp] = scanner.Text()
	}
	scanner.Scan()
	return ArtValue
}

func IsValidArgs(str, sub string) bool {
	if str == "" {
		os.Exit(0)
		return false
	}
	for _, val := range str + sub {
		if val <= 31 || val >= 127 {
			return false
		}
	}
	return true
}

func IsOnlyNewLine(str []string) bool {
	for _, v := range str {
		if v != "" {
			return false
		}
	}
	return true
}


func IsValidColor(color string)bool{

	if color == "" {
		return false
	}
	if !S.HasPrefix(color, "--color="){
		return false
	}
	if len("--color=") == len(color){
		return false
	}
	if !S.Contains("red green yellow blue magenta white", color[len("--color="):]){
		return false
	}
return true
}

