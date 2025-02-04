package main

import (
	"bufio"
	F "fmt"
	"os"
)

var _map = make(map[int][8]string)
var lines = [8]string{}

func Printing(arg string) {
	for _, val := range arg {
		for i := 0 ; i < 8 ; i++{
			lines[i] += _map[int(val)][i] 
		}
	}
	for i := 0; i < 8; i++ {
		F.Println(lines[i])
	}
}

func InitMap() {
	file, err := os.Open("standard.txt")
	if err != nil {
		F.Println("Err: opening file:", err.Error())
		return
	}
	
	scanner := bufio.NewScanner(file)
	
	for i := 32; i < 127; i++ {
		_map[i] = InsertValue(scanner)
	}

	defer file.Close()
}

func InsertValue(scanner *bufio.Scanner) [8]string {
	ArtValue := [8]string{}
	
	for cp := 0; cp < 8 && scanner.Scan(); cp++ {
		ArtValue[cp] = scanner.Text() + " "
	}
	scanner.Scan()
	return ArtValue
}


func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		F.Println("Err: Not Enough Parameters!")
		return
	}

	if len(args) != 1 {
		return
	}

	InitMap()
	Printing(args[0])
}
