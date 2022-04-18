package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(elementCount(doc))
}

func elementCount(n *html.Node) int {
	cnt := 0
	if n.Type == html.ElementNode {
		cnt = 1
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cnt += elementCount(c)
	}
	return cnt
}
