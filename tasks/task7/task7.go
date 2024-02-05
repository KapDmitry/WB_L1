package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.

Для конкурентной записи в map необходимо решить проблему синхронизации доступа нескольких горутин к данным.
Я использовал mutex, для блокировки доступа к данным у других горутин.
Горутина, которая первая "возьмет" mutex и будет писать в map
*/
type ConcMap struct {
	Mu   *sync.Mutex
	Cash map[int]int
}

func NewConcMap() *ConcMap {
	return &ConcMap{
		Mu:   &sync.Mutex{},
		Cash: make(map[int]int),
	}
}

func (i *ConcMap) Get(ID int) (int, error) {
	i.Mu.Lock()
	defer i.Mu.Unlock()
	if val, ok := i.Cash[ID]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("no such order")
}

func (i *ConcMap) Update(ID int, val int) error {
	i.Mu.Lock()
	defer i.Mu.Unlock()
	if _, ok := i.Cash[ID]; ok {
		return fmt.Errorf("order already exists")
	}
	i.Cash[ID] = val
	return nil
}

func main() {
	mp := NewConcMap()

	wg := &sync.WaitGroup{}

	numIters := rand.Intn(10)
	fmt.Println(numIters)

	wg.Add(numIters)
	for i := 0; i < numIters; i++ {
		go func() {
			defer wg.Done()
			mp.Update(rand.Intn(100), rand.Intn(100))
		}()
	}
	wg.Wait()
	fmt.Println(mp.Cash)

}
