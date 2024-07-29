package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)

	if a < 14 {
		fmt.Println("детство")
	} else if a < 25 {
		fmt.Println("молодость")
	} else if a < 60 {
		fmt.Println("зрелость")
	} else {
		fmt.Println("старость")
	}
}
