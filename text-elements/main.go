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
	outputTextElements(doc)
}

func outputTextElements(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outputTextElements(c)
	}
}
