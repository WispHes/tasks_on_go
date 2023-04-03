package main

import "fmt"

func main() {
	i, s := 3, []int{2, 3, 4, 2, 1, 5, 6}
	fmt.Println(elementDeleting(s, i))
}

func elementDeleting(s []int, i int) []int {
	s = append(s[:i], s[i+1:]...)
	return s
}
