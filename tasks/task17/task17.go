package main

import (
	"fmt"
)

/*
binarySearch принимает отсортированный массив и искомый элемент. Она возвращает индекс элемента в массиве, если он найден, и -1, если элемент отсутствует.

Основной цикл бинарного поиска выполняется до тех пор, пока low меньше или равно high. На каждом шаге определяется середина массива (mid), и сравнивается с искомым элементом. В зависимости от результата сужается интервал поиска. Если элемент найден, возвращается его индекс, иначе возвращается -1.
*/

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid // Найден элемент, возвращаем его индекс
		} else if arr[mid] < target {
			low = mid + 1 // Искомый элемент находится справа от mid
		} else {
			high = mid - 1 // Искомый элемент находится слева от mid
		}
	}

	return -1 // Элемент не найден
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 6

	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден по индексу %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}