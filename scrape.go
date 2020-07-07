package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type shoeData struct {
	Name  string
	Price string
}

func printShoes(n []shoeData) {
	fmt.Println("\n-- Nike Mens Running Shoes --")
	for _, item := range n {
		fmt.Printf("Name: %v\nPrice: %v\n\n", item.Name, item.Price)
	}
}

func main() {
	shoes := []shoeData{}

	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("#Wall > div > div.results__body > div > main > section > div", func(e *colly.HTMLElement) {
		e.ForEach(".product-card", func(_ int, e *colly.HTMLElement) {
			item := shoeData{}
			item.Name = e.ChildText(".product-card__info .product-card__title")
			item.Price = e.ChildText(".product-card__info .product-card__price")
			shoes = append(shoes, item)
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting: %v\n", r.URL.String())
	})

	c.Visit("https://www.nike.com/w/mens-shoes-nik1zy7ok")
	printShoes(shoes)

}
