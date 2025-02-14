package main

import (
	"bufio"
	"os"

	A "asciiartcolor/utils"
	F "fmt"
	S "strings"
)

var (
	_map  = make(map[int][8]string)
	lines = [8]string{}
)

const  (
	red = "\033[31m"
	reset = "\033[0m"
)



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
		_map[i] = A.InsertValue(scanner)
	}

	defer file.Close()
}

// func Printing(inp string) {
// 	if inp == "\\n" {
// 		F.Println()
// 		return
// 	}

// 	SplArgs := S.Split(inp, "\\n")

// 	// Fix Multiple "\n"
// 	if A.IsOnlyNewLine(SplArgs) {
// 		for i := 0; i < len(SplArgs)-1; i++ {
// 			F.Println()
// 		}
// 		return
// 	}
// 	// applying "\n"
// 	for _, arg := range SplArgs {
// 		if arg == "" {
// 			F.Println()
// 			continue
// 		}
// 		// Storing Data
// 		for _, val := range arg {
// 			for i := 0; i < 8; i++ {
// 				lines[i] += _map[int(val)][i]
// 			}
// 		}
// 		// Printing Data
// 		for i := 0; i < 8; i++ {
// 			F.Println(lines[i])
// 			lines[i] = ""
// 		}
// 	}
// }



func Printing(str, color, sub string) {

	spl := S.Split(str, sub);
	F.Println(spl)
	
	for _,arg := range spl{
			if arg == ""{
				for _,v:= range sub{
					for i:=0 ; i<8 ; i++{
						lines[i] += red + _map[int(v)][i] + reset 
					}
				}
			}else{
				for _,v:= range arg{
					for i:=0 ; i<8 ; i++{
						lines[i] += _map[int(v)][i] 
					}
				}
			}
		}
		for i:=0 ; i<8 ; i++{
			F.Println(lines[i])
			}
	}
	
	


func main() {
	args := os.Args[1:]

	if (len(args) < 3 || len(args) > 4 ) {
		os.Stderr.WriteString("Usage: go run . --color=<color> <substring to be colored> [STRING]")
		return
	}

	if len(args) == 3 { args = append(args, "standard")}
	
	if !A.IsValidColor(args[0]){
		os.Stderr.WriteString("Err: Invalid Color\n")
		return
	}
	color := S.TrimPrefix(args[0], "--color=")

	if !A.IsValidArgs(args[2], args[1]) {
				os.Stderr.WriteString("Err: Invalid Arguments !\n")
		return
	}

	F.Println(color,"----", args[2],"----",args[1])
	InitMap(args[3]) 
	Printing( args[2], color, args[1])
}
