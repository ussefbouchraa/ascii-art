package main

import (
	"bufio"
	F "fmt"
	"os"
)


var _map = make(map[string]int)

// func IsPrintable(elm string) bool{

// }

func InitMap(){
	_map["H"] = 1
	_map["e"] = 11
	_map["l"] = 20
	_map["o"] = 30
}

func ReadChar(elm rune) string{
	line ,cp := 0, 0
	res := ""
	file, err := os.Open("Banner.txt")
	if err != nil{
		return err.Error()
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		if _map[string(elm)] != line{
			line++
			continue
		}
		break
	}
	if _map[string(elm)] == line {
		for scanner.Scan(){
			if cp == 8{
				break
			}
			res += scanner.Text() + "\n"
			cp++
		}
	}
	return res
}

func main(){
	args:= os.Args[1:]
	if len(args) == 0{
		F.Println("Err: Not Enough Prameters !")
		return 
	}
	if len(args) != 1{
		return
	}
	if args[0] == "" { return }
	if args[0] == "\n" { F.Println(); return }

	
	InitMap()

	for _, val := range(args[0]){
		F.Println(ReadChar(val))
	}
}