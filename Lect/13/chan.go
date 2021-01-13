package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	msg := make(chan string)
	msg2 := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			for {
				time.Sleep(time.Millisecond * 50)
				msg <- fmt.Sprintf("ping:%d", n)
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			for {
				time.Sleep(time.Millisecond * 55)
				msg2 <- fmt.Sprintf("PONG:%d", n)
			}
		}(i)
	}
	nMsg := 0
	for {
		select {
		case out := <-msg:
			nMsg++
			fmt.Printf("%s\n", out)
		case out := <-msg2:
			nMsg++
			fmt.Printf("%s\n", out)
		}
		if nMsg > 20 {
			os.Exit(0)
		}
	}
	wg.Wait()
}
