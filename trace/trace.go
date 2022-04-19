package main

import (
	"fmt"
	"time"
)

func trace(msg string) func() {
	start := time.Now()
	fmt.Printf("enter %s\n", msg)
	return func() {
		fmt.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)
}

func main() {
	bigSlowOperation()
}
