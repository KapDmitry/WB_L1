package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

/*
	Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

В программе создается отдельная горутина, которая проверяет пришел ли сигнал завершения и если не пришел, то отправляет данные в канал. Также создается N worker'ов, которые вычитывают данные из канала
с помощью range (т.е. до тех пор, пока канал не будет закрыт), кроме того я сделал wg.Wait(), для того, чтобы дождаться завершения всех горутин при поступлении сигнала, а также для того, чтобы основная горутина не завершилась раньше времени


*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите количество воркеров в качестве аргумента командной строки.")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Ошибка чтения количества воркеров:", err)
		return
	}

	dataChannel := make(chan string)

	var wg sync.WaitGroup

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		counter := 1
		for {
			select {
			case <-signalChannel:
				close(dataChannel)
				fmt.Println("aaaa")
				return
			default:
				data := fmt.Sprintf("Data %d", counter)
				fmt.Println("bbb")
				dataChannel <- data
				counter++
			}
		}
	}()

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for data := range dataChannel {
				fmt.Printf("Worker %d received: %s\n", workerID, data)
			}
		}(i + 1)
	}
	wg.Wait()

	fmt.Println("Программа завершена.")
}
