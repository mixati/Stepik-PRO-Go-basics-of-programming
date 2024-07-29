package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)

	if math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
