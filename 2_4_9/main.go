package main

import "fmt"

func main() {
	var a, b, n int
	_, _ = fmt.Scan(&a, &b, &n)
	b = (a*100 + b) * n
	fmt.Println(b/100, b%100)
}
