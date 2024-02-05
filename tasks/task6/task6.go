package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.

1. Можно использовать специальный канал по которому мы будем отправлять сигнал, означающий необходимость завершения
2. Можно использовать контекст и завершать горутины при его отмене
3. Можно использовать таймеры (ticker, time.After)
4. Можно убить main
Дополнительные "методы"
5. По сигналу POSIX (по сути, перерождается либо в п1, либо в п2)
6. Можно завершить горутину исходя из выполнения какого-то услвия(значение переменной или что угодно, после выполнения дочерних горутин и т.д.)
*/

func chanStop(ch <-chan int) {
	for {
		select {
		case <-ch:
			fmt.Println("killed by channel")
			return
		default:
			fmt.Println("working hard")
		}
	}
}

func ctxStop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("killed by context")
			return
		default:
			fmt.Println("working hard-2")
		}
	}
}

func timerStop(numSeconds int) {
	tC := time.After(time.Duration(numSeconds) * time.Second)
	for {
		select {
		case <-tC:
			fmt.Println("killed by timer")
			return
		default:
			fmt.Println("working hard-3")
		}
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Укажите количество секунд в качестве аргумента командной строки.")
		return
	}

	numSeconds, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Ошибка чтения количества секунд:", err)
		return
	}

	closeCh := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	go chanStop(closeCh)
	go ctxStop(ctx)
	go timerStop(numSeconds)

	closeCh <- 1 // программа отменит контекст и отправит данные в канал, но может не дождаться завершения горутин
	cancel()

	time.Sleep(time.Duration(numSeconds) * 2 * time.Second) // чтобы точно дождались

}
