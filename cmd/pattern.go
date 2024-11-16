package main

import (
	"context"
	"fmt"
	"time"
)

type data string

func algo(ctx context.Context) chan data {
	dataChan := make(chan data)
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			dataChan <- "MyDATA"
		}
	}()
	return dataChan
}

func UIUpdate(ctx context.Context) {
	dataChan := algo(ctx)

	fmt.Println("MAIN MENU PAUSE")
	// здесь пользователь в main menu, в горутине считается уже первый результат функции algo
	time.Sleep(time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Millisecond * 500):
			// здесь каждый 0.5 секунды мы будем отдавать данные из канала,
			// в горутине после передачи будет начинаться новый подсчёт
			newData := <-dataChan
			fmt.Println(newData)
		}
	}

}

//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	UIUpdate(ctx)
//}
