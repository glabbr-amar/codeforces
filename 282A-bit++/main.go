package main

import "fmt"

func main() {

	var testcase int
	fmt.Scan(&testcase)

	initialValue := 0
	for i := 0; i < testcase; i++ {
		var string string
		fmt.Scan(&string)

		if string[1] == '+' {
			initialValue = initialValue + 1
		} else if string[1] == '-' {
			initialValue = initialValue - 1
		}
	}
	fmt.Println(initialValue)

}
