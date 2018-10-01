package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func getResponse(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get url: %v\n", err)
		os.Exit(1)
	}
	return resp, err
}

func getDoc(err error, resp *http.Response) *html.Node {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get doc: %v\n", err)
		os.Exit(1)
	}
	return doc
}

func visit(links map[string]string, node *html.Node, nodeStart, attrKey string) map[string]string {
	if node.Type == html.ElementNode && node.Data == nodeStart {
		for _, a := range node.Attr {
			if a.Key == attrKey {
				links[a.Val] = ""
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c, nodeStart, attrKey)
	}
	return links
}

func getLinksMap(doc *html.Node, nodeStart, attrKey string) map[string]string {
	links := make(map[string]string)
	links = visit(links, doc, nodeStart, attrKey)
	return links
}

func main() {
	url := os.Args[1] //"https://goods.ru/"
	resp, err := getResponse(url)
	defer resp.Body.Close()

	doc := getDoc(err, resp)
	links := getLinksMap(doc, "img", "src")

	for key := range links {
		fmt.Println(key)
	}
}
