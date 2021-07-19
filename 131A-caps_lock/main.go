package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var str string
	fmt.Scan(&str)
	check := true
	for i := 1; i < len(str); i++ {
		if unicode.IsLower(rune(str[i])) {
			check = false
			break
		}
	}
	ansString := ""
	if check == true {
		for i := 0; i < len(str); i++ {
			if unicode.IsLower(rune(str[i])) {
				ansString += strings.ToUpper(string(str[i]))
			} else {
				ansString += strings.ToLower(string(str[i]))
			}
		}
	}else{
		ansString = str
	}
	fmt.Print(ansString)

}
