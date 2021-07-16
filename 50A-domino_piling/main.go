package main

import "fmt"

func main() {
	var (
		row, column int
	)
	fmt.Scan(&row, &column)
	numberOfDominoes := 0
	numberOfDominoes += (column/2) * row
	if column%2==1 {
		numberOfDominoes += row/2
	}
	fmt.Println(numberOfDominoes)
}
