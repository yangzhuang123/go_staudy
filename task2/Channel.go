package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			channel <- i
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			msg := <-channel
			fmt.Println(msg)
		}
	}()

	time.Sleep(time.Second * 3)

	channel1 := make(chan int, 10)

	go func() {
		for i := 1; i <= 100; i++ {
			channel1 <- i
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			msg := <-channel1
			fmt.Println(msg)
		}
	}()

	time.Sleep(time.Second * 3)

}
