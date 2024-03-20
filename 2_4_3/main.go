package main

import "fmt"

func main() {
	var a int
	_, _ = fmt.Scan(&a)
	var b int = a * a
	var c int = b * b
	fmt.Println(b * c)
}
