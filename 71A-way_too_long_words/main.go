package main

import "fmt"

func main() {
	var numberOfWords int
	fmt.Scanln(&numberOfWords)
	for i := 0; i < numberOfWords; i++ {
		var input string
		fmt.Scanln(&input)
		if len(input) <= 10 {
			fmt.Println(input)
		} else {
			fmt.Print(string(input[0]), len(input)-2, string(input[len(input)-1]))
			fmt.Println()
		}
	}
}
