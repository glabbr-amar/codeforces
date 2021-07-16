package main

import "fmt"

func main() {
	var weight int
	fmt.Scanln(&weight)
	if weight > 2 {
		if weight%2==0 {
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}
	}else{
		fmt.Println("NO")
	}
}
