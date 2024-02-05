package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
В Go множество можно представить с использованием map, где ключами будут уникальные элементы, а значениями - фиктивные значения (например, true или struct{})
В этой задаче createSet принимает массив строк и возвращает множество в виде map[string]bool.
Каждый элемент массива добавляется в множество как ключ мапы с значением true. Таким образом, получается множество строк без повторений.
*/

func createSet(strings []string) map[string]bool {
	set := make(map[string]bool)

	for _, str := range strings {
		set[str] = true
	}

	return set
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	stringSet := createSet(strings)
	fmt.Println("Множество:", stringSet)

	fmt.Println("Присутствие 'dog':", stringSet["dog"])
	fmt.Println("Присутствие 'bird':", stringSet["bird"])
}
