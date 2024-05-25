package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("scraper in go")

	c := colly.NewCollector()
	fmt.Println(c)
}
