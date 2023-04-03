package main

import "fmt"

func main() {
	input := []string{
		"cat",
		"cat",
		"dog",
		"cat",
		"tree",
	}
	fmt.Println(RemoveDuplicates(input))
}

func RemoveDuplicates(input []string) []string {
	m := make(map[string]int)
	res := make([]string, 0, len(input))
	for i, v := range input {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
		m[v] = i
	}
	return res
}
