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
	var links []string
	links = visit(links, doc)
	for _, link := range links {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
