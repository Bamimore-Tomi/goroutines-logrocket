package main

import (
	"fmt"
)

func main() {
	// creating a channel by declaring it
	var mychannel1 chan int
	fmt.Println(mychannel1)

	// creating a channel using make()

	mychannel2 := make(chan int)
	fmt.Println(mychannel2)

}
