package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	line1 := scanner.Text()
	_ = scanner.Scan()
	line2 := scanner.Text()
	_ = scanner.Scan()
	line3 := scanner.Text()

	fmt.Println(line3)
	fmt.Println(line2)
	fmt.Println(line1)
}
