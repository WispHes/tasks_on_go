package main

import "fmt"

func main() {
	fmt.Println(uniqueSymbols("abCdefAaf"))
}

func uniqueSymbols(x string) bool {
	m := make(map[rune]int)
	for _, v := range x {
		m[v]++
		fmt.Println(v)
		if m[v] > 1 {
			return false
		}
	}
	return true
}
