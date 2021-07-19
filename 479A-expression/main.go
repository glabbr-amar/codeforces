package main

import "fmt"

func main() {
	var (
		a, b, c int
	)
	fmt.Scan(&a, &b, &c)
	array := make([]int, 6)
	array[0] = a + b + c
	array[1] = a * b * c
	array[2] = a + (b * c)
	array[3] = a * (b + c)
	array[4] = (a + b) * c
	array[5] = (a * b) + c
	maximum := array[0]
	for i := 1; i < 6; i++ {
		if array[i] > maximum {
			maximum = array[i]
		}
	}
	fmt.Print(maximum)

}
