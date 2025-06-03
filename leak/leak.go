package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var globalFunc func() // Holds reference, causes memory leak

func main() {
	// Start pprof server
	go func() {
		log.Println("Starting pprof server on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Simulate memory leak
	leaky()

	// Keep program running
	select {}
}

func leaky() {
	huge := make([]byte, 100*1024*1024) // 100MB
	fmt.Println("Allocated 100MB")

	// Store closure that captures 'huge' slice
	globalFunc = func() {
		fmt.Println("Using huge data", huge[0])
	}
}
