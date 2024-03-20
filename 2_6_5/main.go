package main

import (
	"fmt"
)

func main() {
	var a float64
	fmt.Scan(&a)
	fmt.Println(a - float64(int(a)))
}
