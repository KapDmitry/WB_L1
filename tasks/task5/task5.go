package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

Для выполения условия задания я решил воспользоваться time.After, который создает канал, и записывает туда данные спустя заданное время. Таким образом, есть горутина, которая публикует данные в канал, но когда
в канал time.After приходит уведомление она прекращает выполнение и завершает публикацию данных, закрывает канал.
В основной же горутине происходит считывание из канала с данными и оно закончится, когда он будет закрыт
*/

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

	tC := time.After(time.Duration(numSeconds) * time.Second)
	dataChannel := make(chan int)
	go func() {
		for {
			select {
			case <-tC:
				close(dataChannel)
				return
			default:
				dataChannel <- 1
			}
		}
	}()

	for v := range dataChannel {
		fmt.Println(v)
	}

	fmt.Println("Программа завершена")

}
