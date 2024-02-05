package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
1. Проверяем каждый элемент первого множества
2. Если элемент присутствует во втором множестве, добавляем его в результирующее множество(пересечение)
*/

func intersection(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)

	for elem := range set1 {
		if set2[elem] {
			result[elem] = true
		}
	}

	return result
}

func main() {
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true, 7: true}

	intersectedSet := intersection(set1, set2)

	fmt.Println("Пересечение:", intersectedSet)
}
