package main

import (
	"fmt"
	"sync"
)

/*
Counter - это структура с полем value и мьютексом mu. Методы Increment и GetValue защищены мьютексом для обеспечения безопасного доступа к общим данным в конкурентной среде.

Программа создает несколько горутин, каждая из которых инкрементирует счетчик 100000 раз. После завершения всех горутин программа выводит итоговое значение счетчика.
*/

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	finalValue := counter.GetValue()
	fmt.Println("Final Counter Value:", finalValue)
}
