package main

import (
	"encoding/csv"
	"log"
	"os"

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

	file, err := os.Create("books.csv")
	if err != nil {
		log.Fatalln("failed to open", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}
	writer.Write(headers)

	for _, book := range books {
		rec := []string{
			book.url,
			book.image,
			book.name,
			book.price,
		}
		writer.Write(rec)
	}
	defer writer.Flush()
}
