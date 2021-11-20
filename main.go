package main

import (
	"fmt"
	"math/rand"
	"time"
)

// name is an string to identify the function call
// limit the amount of number the funcion will print
// sleep is the amount of seconds before the function prints the next value
func randSleep(name string, limit int, sleep int) {
	for i := 1; i < limit; i++ {
		fmt.Println(name, rand.Intn(i))
		time.Sleep(time.Duration(sleep * int(time.Second)))

	}

}
func main() {
	go randSleep("first:", 4, 3)
	go randSleep("second:", 4, 3)

}

// git checkout 01
