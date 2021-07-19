package main

import "fmt"

func main() {
	var stringSize int
	var str string
	fmt.Scan(&stringSize,&str)
	ansString := ""
	currentChar := str[0]
	ansString += string(currentChar)
	for i := 1; i < stringSize; i++ {
		if str[i] != currentChar {
			currentChar = str[i]
			ansString += string(currentChar)
		}
	}
	fmt.Print(len(str) - len(ansString))
}
