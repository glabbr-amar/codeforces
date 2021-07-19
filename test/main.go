// 230A-dragons
package main

import (
	"fmt"
	"sort"
)

func sortByKey(mp map[int]int) []int {
	var keys []int
	for k, _ := range mp {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func canDefeat(numberOfDragon int,strength int,dragonStrengthArray []int,mp map[int]int) bool {
	cnt:=0
	for i := 0; i < numberOfDragon; i++ {
		if strength >= dragonStrengthArray[i] {
			strength += mp[dragonStrengthArray[i]]
			cnt++
		}else{
			break
		}
	}
	if cnt==numberOfDragon{
		return true
	}else{
		return false
	}
}

func main() {
	var (
		strength, numberOfDragon, dragonStrength, bonusStrength int
	)
	fmt.Scan(&strength, &numberOfDragon)
	mp := make(map[int]int)

	for i := 0; i < numberOfDragon; i++ {
		fmt.Scan(&dragonStrength, &bonusStrength)
		mp[dragonStrength] = bonusStrength
	}


	dragonStrengthArray := sortByKey(mp)
	canDefeat := canDefeat(numberOfDragon,strength,dragonStrengthArray,mp)

	if canDefeat {
		fmt.Print("YES")
	}else{
		fmt.Print("NO")
	}

}
