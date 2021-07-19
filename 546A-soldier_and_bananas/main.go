package main

import "fmt"

func main() {
	var(
		costOfFirstBanana int64
		initialMoney int64
		numberOfBanana int64
	)
	fmt.Scan(&costOfFirstBanana, &initialMoney, &numberOfBanana)
	requiredMoney := costOfFirstBanana * numberOfBanana * (numberOfBanana+1) / 2
	if initialMoney<requiredMoney {
		fmt.Print(requiredMoney - initialMoney)
	} else{
		fmt.Print(0)
	}
}

