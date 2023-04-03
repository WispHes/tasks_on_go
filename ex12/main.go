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
	// Первый вариант решения без сохранения последовательности
	fmt.Println(RemoveDuplicates1(input))
	// Второй вариант решения с сохраненем последовательности
	fmt.Println(RemoveDuplicates2(input))
}

func RemoveDuplicates1(items []string) []string {
	m := make(map[string]int)
	for _, v := range items {
		m[v]++
	}
	/* здесь можно было вернуть сразу map
	но так как я решил сделать независимыми два решения
	пришлось результат поместить в слайс */
	res := make([]string, 0, len(m))
	for key := range m {
		res = append(res, key)
	}
	return res
}

func RemoveDuplicates2(items []string) []string {
	m := make(map[string]int)
	res := make([]string, 0, len(items))
	for i, v := range items {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
		m[v] = i
	}
	return res
}
