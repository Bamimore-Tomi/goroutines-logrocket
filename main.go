package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg is the pointer to a waitgroup
// name is an string to identify the function call
// limit the amount of number the funcion will print
// sleep is the amount of seconds before the function prints the next value
func randSleep(wg *sync.WaitGroup, name string, limit int, sleep int) {
	defer wg.Done()
	for i := 1; i <= limit; i++ {
		fmt.Println(name, rand.Intn(i))
		time.Sleep(time.Duration(sleep * int(time.Second)))

	}

}
func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go randSleep(wg, "first:", 10, 2)
	go randSleep(wg, "second:", 3, 2)
	wg.Wait()

}

// OUTPUT

// second: 0
// first: 0
// first: 1
// second: 1
// second: 1
// first: 0
// first: 1
// first: 0
// first: 4
// first: 1
// first: 6
// first: 7
// first: 2

// git checkout 03
