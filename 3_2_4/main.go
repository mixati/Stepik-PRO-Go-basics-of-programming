package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)

	tmp1 := a / 100
	tmp2 := a/10 - tmp1*10
	tmp3 := a % 10
	if tmp1 != tmp2 && tmp1 != tmp3 && tmp2 != tmp3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
