package asciiArt

import "bufio"

func InsertValue(scanner *bufio.Scanner) [8]string {
	ArtValue := [8]string{}
	
	for cp := 0; cp < 8 && scanner.Scan() ; cp++ {
		ArtValue[cp] = scanner.Text() 
	}
	scanner.Scan()
	return ArtValue
}

func IsValidArg(arg string) bool{
	for _,val := range arg{
		if val <= 31 || val >=127 {
			return false
		}
	}
	return true
}


func IsOnly(arg []string) bool {
	for _,val := range arg{
		if string(val) != "" {
			return false
		}
	}
	return true 
}