package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int
var url string = "https://www.simpleweb.org/"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		resp.Body.Close()
		log.Fatalln("get: ", url, "falied.")
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalln("get: ", url, " status: ", resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}
	forEachNode(doc, startElement, shortFormElement, endElement)
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func shortFormElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s />\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func forEachNode(n *html.Node, pre, inCondition, post func(n *html.Node)) {
	hasChild := n.FirstChild != nil && inCondition != nil
	if !hasChild {
		inCondition(n)
	}
	if pre != nil && hasChild {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, inCondition, post)
	}
	if post != nil && hasChild {
		post(n)
	}
}
