package main

import (
	"fmt"
	"golang-demos/bytecounter/bytecounter"
)

func main() {
	var c bytecounter.ByteCounter
	fmt.Fprintf(&c, "%d", 10)
	fmt.Println(c)
}
