package main

import (
	"fmt"
	"sync"
)

func writeChannel(wg *sync.WaitGroup, limitchannel chan int, stop int) {
	defer wg.Done()
	for i := 1; i <= stop; i++ {
		limitchannel <- i
	}

}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	limitchannel := make(chan int, 2)
	defer close(limitchannel)
	go writeChannel(wg, limitchannel, 2)
	wg.Wait()
	fmt.Println(<-limitchannel)
	fmt.Println(<-limitchannel)

}

// OUTPUT

// 1
// 2

// git checkout 05
