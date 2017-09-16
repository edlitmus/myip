package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("http://checkip.dyndns.org")
	if err != nil {
		log.Fatal(err)
	}

	htm, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	parse(htm)
	resp.Body.Close()
}

func parse(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "body" {
		fmt.Print(node.FirstChild.Data)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		parse(c)
	}
}
