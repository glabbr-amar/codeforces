package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	str = strings.ToLower(str)
	ansString := ""
	for _, val := range str {
		if val != 'a' && val != 'A' && val != 'e' && val != 'E' && val != 'i' && val != 'I' && val != 'o' && val != 'O' && val != 'u' && val != 'U' && val != 'y' && val != 'Y' {
			ansString += "."
			ansString += string(val)
		}
	}
	fmt.Print(ansString)
}
