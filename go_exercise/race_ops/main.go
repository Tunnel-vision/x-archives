package main

import (
	"fmt"
	"runtime"
	"time"
)

var i = 0

func main() {
	runtime.GOMAXPROCS(2)

	go func() {
		for {
			fmt.Println("i is", i)
			time.Sleep(time.Second)
		}
	}()

	for {
		i += 1
	}
}
