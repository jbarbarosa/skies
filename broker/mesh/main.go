package main

import (
	"fmt"
	"time"
)

func f(from string, number int) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, number, ":", i)
	}
}

func main() {

	f("direct", 0)

	for i := 1; i < 100; i++ {
		go f("goroutine", i)
	}

	time.Sleep(time.Second)
	fmt.Println("done")
}

