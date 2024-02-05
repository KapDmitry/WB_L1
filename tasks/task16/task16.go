package main

import "fmt"

/*
QuickSort
Находим разделение - выбираем pivot из середины массива и меняем местами его и последний элемент, чтобы точно знать индекс для замены после того, как загоню все элементы, меньшие его, налево
Далее перебрасываем все элементы, меньшие пивота, влево и двигаем указатель left
По итогу он будет указывать на первый элемент, который >= pivot
Я бы хотел, в дальнейшем обрабатывать два массива, один с элементами меньше пивота, другой с элементами >= пивота
Для этого я делаю nums[left] = nums[r], то есть ставлю пивот на место разделяющего элемента
и запускаю дальнейшую рекурсию с параметрами left-1, left + 1, обрабатывая оба подмассива без пивота

*/
func quickSort(nums []int, l int, r int) {
	if l < r {
		pivotIndex := partition(nums, l, r)
		quickSort(nums, l, pivotIndex-1)
		quickSort(nums, pivotIndex+1, r)
	}
}

func partition(nums []int, l int, r int) int {
	pivot := nums[(l+r)/2]
	left := l
	nums[(l+r)/2], nums[r] = nums[r], nums[(l+r)/2]

	// перебрасываем все элементы меньше пивота налево
	for i := l; i < r; i++ {
		if nums[i] < pivot {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	nums[left], nums[r] = nums[r], nums[left]
	return left
}

func main() {
	//arr := []int{7, 2, 1, 6, 8, 5, 3, 4}
	arr := []int{3, 7, 8, 5, 2, 1, 9, 5, 4}
	fmt.Println("Unsorted array:", arr)

	// Вызов QuickSort
	quickSort(arr, 0, len(arr)-1)

	fmt.Println("Sorted array:", arr)
}
