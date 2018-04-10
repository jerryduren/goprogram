package main

import (
	"fmt"
	"golang.org/x/net/html"
	"math"
	"os"
)

func Hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	fmt.Printf("%T\n", Hypot) //Print the type of variable Hypot
	fmt.Println(Hypot(30, 40))

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
