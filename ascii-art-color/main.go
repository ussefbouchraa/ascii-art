package main

import (
	"bufio"
	"os"
	A "asciiartcolor/utils"
	F "fmt"
	S "strings"
)

var (
	asciiMap  = make(map[int][8]string)
	lines = [8]string{}
	colorMap = map[string]string{
        "red":     "\033[31m",
        "green":   "\033[32m",
        "yellow":  "\033[33m",
        "blue":    "\033[34m",
        "magenta": "\033[35m",
        "white":   "\033[37m",
        "reset":   "\033[0m",
    }
)

func printlines(){
	for i:=0 ; i < 8 ; i++{
		F.Println(lines[i])
		lines[i] = ""
	}
}

func InitMap(banner string) {
	var file *os.File
	var err error

	switch banner {
		
	case "standard", "standard.txt":
		file, err = os.Open("Banners/standard.txt")
	case "shadow", "shadow.txt":
		file, err = os.Open("Banners/shadow.txt")
	case "thinkertoy", "thinkertoy.txt":
		file, err = os.Open("Banners/thinkertoy.txt")
	default:
		os.Stderr.WriteString("Err: Invalid Argument [BANNER]: " + banner + "\n")
		os.Exit(0)
	}

	if err != nil {
		os.Stderr.WriteString("Err: " + err.Error() + "\n")
		os.Exit(0)
	}

	scanner := bufio.NewScanner(file)
	// to skip first empty line
	if banner == "shadow" || banner == "thinkertoy" ||
		banner == "shadow.txt" || banner == "thinkertoy.txt" {
		scanner.Scan()
	}
	// insert data on the map
	for i := 32; i < 127; i++ {
		asciiMap[i] = A.InsertValue(scanner)
	}

	defer file.Close()
}


func Printing(str, color, sub string) {
	if str == "\\n" { F.Println(); return }

	SplArgs := S.Split(str, "\\n")
	if A.IsOnlyNewLine(SplArgs) {
		for i := 0; i < len(SplArgs)-1; i++ {
			F.Println()
		}
		return
	}
	for _,args := range SplArgs{
		if args == "" {F.Println(); continue}

	args = S.Replace(args, sub, "\x00", -1)
	for _,arg := range args{
		if string(arg) == "\x00"{
			for _,v:= range sub{
				for i:=0 ; i<8 ; i++{
					lines[i] += colorMap[color]+ asciiMap[int(v)][i] + colorMap["reset"] 
				}
			}
		}else{
			for _,v:= range string(arg){
				for i:=0 ; i<8 ; i++{
					lines[i] += asciiMap[int(v)][i] 
				}
			}
		}
	}
		printlines()
}
}
	
	
func main() {
	args := os.Args[1:]


if len(args) == 1 || len(args) == 2{
	if len(args) == 1{	args = append(args, "standard")}
	if !A.IsValidArgs(args[0], "") {
		os.Stderr.WriteString("Err: Invalid Arguments [STRING]!\n")
		return
	}
	InitMap(args[1])
	Printing( args[0], "", "") 
	return
}

	if (len(args) < 3 || len(args) > 4 ) {
		os.Stderr.WriteString("Usage: go run . --color=<color> <substring to be colored> [STRING]\n")
		return
	}

	if len(args) == 3 { args = append(args, "standard")}
	
	if !A.IsValidColor(args[0]){
		os.Stderr.WriteString("Usage: go run . --color=<color> <substring to be colored> [STRING]\n")
		return
	}

	if !A.IsValidArgs(args[2], args[1]) {
			os.Stderr.WriteString("Err: Invalid Arguments [STRING] [SUB] \n")
		return
	}

	InitMap(args[3])  //banner
	Printing( args[2], args[0][8:], args[1]) // string  || color || sub 
}
