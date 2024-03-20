package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(float64(a) / math.Pow(2, 13))
}
