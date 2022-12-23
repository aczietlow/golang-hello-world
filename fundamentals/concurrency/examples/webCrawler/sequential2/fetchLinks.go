package main

import (
	"fmt"
	"github.com/aczietlow/hello/fundamentals/functions/links"
	"log"
)

func main() {
	// list of urls
	var urls = make([]string, 0)

	// provide init urls.
	for _, v := range []string{"https://zietlow.io/"} {
		urls = append(urls, v)
	}

	// crawl urls.
	breadthFirst(crawl, urls)

}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Fetch(url)
	if err != nil {
		log.Println(err)
	}
	return list
}
