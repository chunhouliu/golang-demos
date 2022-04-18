package main

import (
	"fmt"
	"golang-demos/links/links"
	"log"
	"os"
)

func bfs(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	urls, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return urls
}

func main() {
	bfs(crawl, os.Args[1:])
}
