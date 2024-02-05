package main

import (
	"fmt"
	"time"
)

func mySleep(seconds int) {
	startTime := time.Now()
	targetTime := startTime.Add(time.Duration(seconds) * time.Second)

	for time.Now().Before(targetTime) {
	}

	return
}

func main() {
	fmt.Println("Before sleep")

	mySleep(3)

	fmt.Println("After sleep")
}
