package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)

	b := a / 1000
	b1 := b / 100
	b2 := b/10 - b1*10
	b3 := b%10 + b1 + b2

	c := a % 1000
	c1 := c / 100
	c2 := c/10 - c1*10
	c3 := c%10 + c1 + c2

	if b3 == c3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
