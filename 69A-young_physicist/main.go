package main

import "fmt"

func main() {
	var (
		numberOfInput int
		xCordinate    int
		yCordinate    int
		zCordinate    int
	)
	xSum, ySum, zSum := 0, 0, 0

	fmt.Scan(&numberOfInput)
	for i := 0; i < numberOfInput; i++ {
		fmt.Scan(&xCordinate, &yCordinate, &zCordinate)
		xSum += xCordinate
		ySum += yCordinate
		zSum += zCordinate
	}
	if xSum==0 && ySum==0 && zSum==0 {
		fmt.Print("YES")
	}else{
		fmt.Print("NO")
	}
}
