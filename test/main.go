package main

import "fmt"

func main() {
	str := "amarkumar"
	mp := make(map[byte]int)
	for _,val := range str {
		mp[byte(val)]++
	}
	fmt.Print(len(mp))
}
