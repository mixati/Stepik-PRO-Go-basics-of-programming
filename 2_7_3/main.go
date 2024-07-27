package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a%10 + a%100/10 + a%1000/100)
}
