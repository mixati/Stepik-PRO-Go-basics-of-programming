package main

import "fmt"

func main() {
	var a int
	_, _ = fmt.Scan(&a)
	fmt.Println(a % 10)
}
