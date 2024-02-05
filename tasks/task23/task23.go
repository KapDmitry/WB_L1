package main

import "fmt"

func removeElement(slice []int, index int, needOrder bool) ([]int, error) {
	if index < 0 || index >= len(slice) {

		return slice, fmt.Errorf("out of range")
	}

	if needOrder {
		return append(slice[:index], slice[index+1:]...), nil
	}
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1], nil
}

func main() {
	// Пример использования
	slice := []int{1, 2, 3, 4, 5}
	indexToRemove := 2

	fmt.Printf("Original slice: %v\n", slice)

	// Удаляем элемент по указанному индексу
	slice, err := removeElement(slice, indexToRemove, true)
	if err == nil {
		fmt.Printf(err.Error())
	}

	fmt.Printf("Slice after removing element at index %d: %v\n", indexToRemove, slice)
}
