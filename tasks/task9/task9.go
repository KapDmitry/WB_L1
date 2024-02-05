package main

import (
	"fmt"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

Первая горутина записывает числа в массиве в первый канал (inputCh)
Вторая горутина умножает числа на 2 и записывает результат во второй канал (outputCh)
Main горутина выводит результаты из второго канала в stdout.
*/
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	inputCh := make(chan int)
	outputCh := make(chan int)

	go func() {
		defer close(inputCh)
		for _, num := range numbers {
			inputCh <- num
		}
	}()

	go func() {
		defer close(outputCh)
		for num := range inputCh {
			result := num * 2
			outputCh <- result
		}
	}()

	for result := range outputCh {
		fmt.Println(result)
	}

}
