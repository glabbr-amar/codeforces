package main

import "fmt"

func main() {
	var (
		length  int64
		breadth int64
		side    int64
	)
	fmt.Scan(&length, &breadth, &side)
	row := length / side
	column := breadth / side
	if length%side != 0 {
		row++
	}
	if breadth%side != 0 {
		column++
	}
	fmt.Print(row * column)
}
