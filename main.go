package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type ArtProduct struct {
	url   string
	image string
	name  string
	price string
}

func main() {
	var arts []ArtProduct
	c := colly.NewCollector()

	c.OnHTML(".product", func(e *colly.HTMLElement) {
		ArtProd := ArtProduct{}
		ArtProd.url = e.ChildAttr("a", "href")
		ArtProd.image = e.ChildAttr("img", "src")
		ArtProd.name = e.ChildText(".product b")
		ArtProd.price = e.ChildText(".product i")
		arts = append(arts, ArtProd)
	})
	c.Visit("http://www.pollynorstore.co.uk/products")

	file, err := os.Create("arts.csv")
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

	for _, book := range arts {
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
