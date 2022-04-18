package main

import (
	"fmt"
	"golang-demos/links/links"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		urls, err := links.Extract(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, url := range urls {
			fmt.Println(url)
		}
	}
}
