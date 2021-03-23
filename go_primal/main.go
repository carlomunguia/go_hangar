package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func boring2() {
	c := boring("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow")
			return
		}
	}
}
