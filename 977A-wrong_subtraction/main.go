package main

import "fmt"

func main() {
	var (
		number int
		k      int
	)
	fmt.Scan(&number, &k)
	for i := 0; i < k; i++ {
		if number%10 == 0 {
			number = number / 10
		} else {
			number = number - 1
		}
	}
	fmt.Print(number)
}
