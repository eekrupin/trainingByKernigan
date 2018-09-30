package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {

	url := os.Args[1] //"https://goods.ru/"
	resp, err := getResponse(url)
	defer resp.Body.Close()

	doc := getDoc(err, resp)
	links := getLinksMap(doc)

	for key := range links {
		fmt.Println(key)
	}

}

func getLinksMap(doc *html.Node) map[string]string {
	links := make(map[string]string)
	visit(links, doc)
	return links
}

func getDoc(err error, resp *http.Response) *html.Node {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findImgLinks: %v\n", err)
		os.Exit(1)
	}
	return doc
}

func getResponse(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	return resp, err
}

func visit(links map[string]string, node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "img" {
		for _, a := range node.Attr {
			if a.Key == "src" {
				links[a.Val] = ""
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		visit(links, c)
	}
}
