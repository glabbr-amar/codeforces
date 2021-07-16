package main

import "fmt"

func main() {
	var (
		numberOfParticipants int
		kthPlace             int
	)
	fmt.Scan(&numberOfParticipants, &kthPlace)
	var array = make([]int, numberOfParticipants)
	for i := 0; i < numberOfParticipants; i++ {
		fmt.Scan(&array[i])
	}
	kthPlaceValue := array[kthPlace-1]
	cnt := 0
	for _, val := range array {
		if val > 0 && val >= kthPlaceValue {
			cnt++
		}
	}
	fmt.Println(cnt)
}
