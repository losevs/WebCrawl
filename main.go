package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)
	c.OnHTML(".mw-page-container-inner", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		fmt.Println(links)
	})
	c.Visit("https://en.wikipedia.org/wiki/Go_(programming_language)")
}
