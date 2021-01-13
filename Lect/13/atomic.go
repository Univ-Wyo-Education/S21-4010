package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	state := make(map[int]int)
	mutex := &sync.Mutex{}

	var nRead uint64
	var nWrite uint64

	const randRange = 15

	for ii := 0; ii < 100; ii++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(randRange)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&nRead, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for jj := 0; jj < 50; jj++ {
		go func() {
			for {
				key := rand.Intn(randRange)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&nWrite, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second * 1)

	nReadTotal := atomic.LoadUint64(&nRead)
	nWriteTotal := atomic.LoadUint64(&nWrite)

	mutex.Lock()
	fmt.Printf("ReadOps: %d\nWriteOps: %d\nFinal State: %+v\n", nReadTotal, nWriteTotal, state)
	mutex.Unlock()

}
