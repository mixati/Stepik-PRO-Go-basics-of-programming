package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)

	tmp := a / 100
	if tmp > 0 && tmp < 10 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
