package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)

	if a >= 7 || a <= -3 {
		fmt.Println("Принадлежит")
	} else {
		fmt.Println("Не принадлежит")
	}
}
