package main

import (
	"fmt"
)

func main() {
	var l1, l2 string

	fmt.Scan(&l1, &l2)
	if l1 == l2 {
		fmt.Println("Пароль принят")
	} else {
		fmt.Println("Пароль не принят")
	}
}
