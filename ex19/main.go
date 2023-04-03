package main

import "fmt"

func main() {
	fmt.Println(flippingLine("главрыба"))
}

func flippingLine(items string) string {
	s := []rune(items)
	n := len(s) - 1
	for i := 0; i <= n/2; i++ {
		s[i], s[n-i] = s[n-i], s[i]
	}
	return string(s)
}
