package main

import "fmt"

func main() {
	var (
		limak int
		bob   int
	)
	fmt.Scan(&limak, &bob)
	cnt := 0
	for limak <= bob {
		limak = limak * 3
		bob = bob * 2
		cnt++
	}
	fmt.Print(cnt)
}
