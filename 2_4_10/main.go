package main

import "fmt"

func main() {
	var a int
	_, _ = fmt.Scan(&a)
	fmt.Println(a, "мин - это", a/60, "час", a%60, "минут.")
}
