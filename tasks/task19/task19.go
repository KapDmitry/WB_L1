package main

import (
	"fmt"
)

/*
Преобразуем все в массив run, а потом метод двух указателей
*/

func reverseString(input string) string {
	runes := []rune(input)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	input := "главрыба"
	reversed := reverseString(input)

	fmt.Printf("Original string: %s\n", input)
	fmt.Printf("Reversed string: %s\n", reversed)
}
