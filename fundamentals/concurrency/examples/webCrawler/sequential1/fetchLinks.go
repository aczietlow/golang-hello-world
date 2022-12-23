package main

import (
	"fmt"
	html "golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	// list of urls
	var urls = make([]string, 0)

	// provide init urls.
	for _, v := range []string{"https://zietlow.io/"} {
		urls = append(urls, v)
	}

	for _, v := range urls {
		links, err := fetch(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
		fmt.Println(links)

	}

}

func fetch(url string) ([]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, response.Status)
	}

	htmlDoc, err := html.Parse(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return extractLinks(nil, htmlDoc), nil
}

func extractLinks(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		links = extractLinks(links, child)
	}
	return links
}
