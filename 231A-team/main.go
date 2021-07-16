package main

import "fmt"

func main() {
	var numberOfProblems int
	fmt.Scan(&numberOfProblems)
	sum:=0
	for i:=0 ; i<numberOfProblems ; i++ {
		var (
			petya int
			vasya int
			tonya int
		)
		fmt.Scan(&petya,&vasya,&tonya)
		if petya+vasya+tonya >=2 {
			sum += 1
		}
	}
	fmt.Println(sum)
}

