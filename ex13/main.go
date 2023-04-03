package main

import "fmt"

func main() {
	a, b := 10, 15
	b = a + b
	a = b - a
	b = b - a
	fmt.Println(a, b)
}
