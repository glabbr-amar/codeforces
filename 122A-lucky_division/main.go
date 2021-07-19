package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	luckyNumber := [9]int {4,7,47,74,447,474,477,747,774}
	flag := 0
	for i:=0 ; i<9 ; i++{
		if number%luckyNumber[i] ==0 {
			flag = 1
			break
		}
	}
	if flag==1 {
		fmt.Print("YES")
	}else{
		fmt.Print("NO")
	}
}
