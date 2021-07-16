package main

import (
	"fmt"
	"sort"
)

func main() {
	var input string
	fmt.Scan(&input)

	var array = make([]byte, len(input)/2+1)
	k := 0
	for i := 0; i < len(input); i++ {
		if input[i] != '+' {
			array[k] = input[i]
			k++
		}
	}
	sort.Slice(array, func(i int, j int) bool { return array[i] < array[j] })
	for i := 0; i < len(array); i++ {
		fmt.Print(string(array[i]))
		if i != len(array)-1 {
			fmt.Print("+")
		}
	}

}
