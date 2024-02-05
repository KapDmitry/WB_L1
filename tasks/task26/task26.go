package main

import "fmt"

/*
программа использует мапу (charCount), чтобы отслеживать, был ли символ уже встречен.
Функция toLowerCase приводит символы к нижнему регистру перед проверкой уникальности. В main приведены примеры использования и вывод результата для трех строк.
*/

func areAllCharactersUnique(str string) bool {
	charCount := make(map[rune]bool)

	for _, char := range str {
		lowerCaseChar := toLowerCase(char)

		if charCount[lowerCaseChar] {
			return false
		}

		charCount[lowerCaseChar] = true
	}

	return true
}

func toLowerCase(char rune) rune {
	if 'A' <= char && char <= 'Z' {
		return char + ('a' - 'A')
	}
	return char
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	result1 := areAllCharactersUnique(str1)
	result2 := areAllCharactersUnique(str2)
	result3 := areAllCharactersUnique(str3)

	fmt.Printf("%s: %t\n", str1, result1)
	fmt.Printf("%s: %t\n", str2, result2)
	fmt.Printf("%s: %t\n", str3, result3)
}
