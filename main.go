package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2) // Buffered channel with size 2

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("sending... ", i)
			c <- i // Send value to channel
			fmt.Println("=> sent ", i)
		}
		close(c) // prevent deadlock, the receiver ok == false
	}()

	fmt.Println("pull ", <-c)

	time.Sleep(100 * time.Millisecond)
	// for {
	// 	select {
	// 	case msg, ok := <-c:
	// 		if !ok {
	// 			fmt.Println("Worker received done signal, stopping...")
	// 			return
	// 		}
	// 		fmt.Print("pull ", msg)
	// 		fmt.Print("\n")
	// 	}
	// }
}
