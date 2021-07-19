package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	/*ans :=0
	cnt:=0
	for i:=0 ; i<26 ; i++{
		cnt = strings.Count(str, string(97+i))
		if cnt>0{
			ans+=1
		}
	}
	if ans&1==0 {
		fmt.Print("CHAT WITH HER!")
	} else{
		fmt.Print("IGNORE HIM!")
	}*/
	mp := make(map[byte]int)
	for _, val := range str {
		mp[byte(val)]++
	}
	if len(mp)&1 == 0 {
		fmt.Print("CHAT WITH HER!")
	}else{
		fmt.Print("IGNORE HIM!")
	}
}
