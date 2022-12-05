package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}
	res := outline(nil, doc)
	fmt.Println(res)
}

func outline(stack []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		stack = outline(stack, c)
	}
	return stack
}
