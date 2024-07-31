package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	if a >= -3 && a <= 1 || a >= 5 && a <= 9 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
