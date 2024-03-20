package main

import "fmt"

func main() {
	var a int
	_, _ = fmt.Scan(&a)
	fmt.Println(a%100/10*10 + a%10*100 + a/100)
}
