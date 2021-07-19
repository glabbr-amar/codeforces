package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	hello := [5]byte{'h', 'e', 'l', 'l', 'o'}
	index, flag := 0, 0
	currentChar := hello[index]
	for _, val := range str {
		if byte(val) == currentChar {
			index++
			if index == 5 {
				flag = 1
				break
			}
			currentChar = hello[index]
		}
	}
	if flag == 1 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
