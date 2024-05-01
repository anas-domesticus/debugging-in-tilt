package main

import (
	"fmt"
	"time"
)

func main() {
	testStr := "foo"
	for {
		fmt.Println("I'm alive...")
		time.Sleep(2 * time.Second)
		fmt.Println(testStr)
	}
}
