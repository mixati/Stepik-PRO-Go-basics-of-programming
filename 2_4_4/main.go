package main

import "fmt"

func main() {
	var a, b, c int
	_, _ = fmt.Scan(&a)
	_, _ = fmt.Scan(&b)
	_, _ = fmt.Scan(&c)
	fmt.Println(a * b * c)
}
