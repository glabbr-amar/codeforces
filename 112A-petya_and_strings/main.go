package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		string1 string
		string2 string
	)
	fmt.Scan(&string1, &string2)
	string1 = strings.ToLower(string1)
	string2 = strings.ToLower(string2)
	if string1<string2 {
		fmt.Println("-1")
	}else if string1>string2 {
		fmt.Println("1")
	}else if string1==string2{
		fmt.Println("0")
	}
}
