package main

import "fmt"

func main() {
	var a, b, c, d int
	_, _ = fmt.Scan(&a, &b, &c, &d)
	fmt.Println((a + b + c + d) * 3)
}
