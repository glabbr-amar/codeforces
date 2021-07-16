package main

import (
	"fmt"
	"math"
)

func main() {
	var arr[7][7] int
	var (
		xCordinate, yCordinate int
	)
	for i:=1 ; i<6 ; i++ {
		for j:=1 ; j<6 ; j++ {
			fmt.Scan(&arr[i][j])
			if arr[i][j]==1 {
				xCordinate = i
				yCordinate = j
			}
		}
	}
	answer := math.Abs(float64(3-xCordinate)) + math.Abs(float64(3-yCordinate))
	fmt.Println(answer)
}
