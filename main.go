package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type BookProduct struct {
	url   string
	image string
	name  string
	price string
}

func main() {
	var books []BookProduct
	c := colly.NewCollector()

	c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
		bookProd := BookProduct{}
		bookProd.url = e.ChildAttr("a", "href")
		bookProd.image = e.ChildAttr("img", "src")
		bookProd.name = e.ChildText("h3")
		bookProd.price = e.ChildText(".price_color")
		books = append(books, bookProd)
	})
	c.Visit("http://books.toscrape.com/")

	fmt.Println(books)

}
