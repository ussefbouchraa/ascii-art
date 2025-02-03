package main

import (
	"bufio"
	F "fmt"
	"os"
)


var _map = make(map[string]int)
var lines [8]string


func InitMap(){
	_map["H"] = 1
	_map["e"] = 10
	_map["l"] = 20
	_map["o"] = 30
}

func Printing(){
	for _, val := range(lines){
		F.Println(val)
	}
}

func ReadChar(elm rune) [8]string{
	line ,cp := 0, 0
	file, err := os.Open("Banner.txt")
	if err != nil{
		return [8]string{err.Error()}
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
		for scanner.Scan() && cp < 8 {
			lines[cp] += scanner.Text()	+ " "	
			cp++
		}
	}	
	return lines
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
	if args[0] == "\n" { return }
	
	InitMap()
	
	for _, val := range(args[0]){
		ReadChar( val)
	}
	
	Printing()

}