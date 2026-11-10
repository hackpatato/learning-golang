package main

import (
	"golang.org/x/net/html"
	"net/http"
	"fmt"
)


func Htmlparse() {
	if resp, err := http.Get("test.com"); err == nil {
		if rootNode, err := html.Parse(resp.Body); err == nil {
			for _, link := range FindLinks(rootNode) {
				fmt.Printf("link : %s\n", link)
			}			
		}
	}

}

func FindLinks(n *html.Node) (links []string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}
	for c := n.FristChild; c != nil; c = c.NextSibling {
		links = append(links, FindLinks(c)...)
	}
	return links
}

