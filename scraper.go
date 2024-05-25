package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// collector manages the network communication and responsible for the execution of the attached callbacks while a collector job is running
	c := colly.NewCollector()

	// callbacks in order
	// 1. called before a request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting: ", r.URL)
	})

	// 2. called if error occured during the request
	c.OnError(func(r *colly.Response, err error) {
		// log.Println("something went wrong: ",err)
		fmt.Println("failed with response: ", r, "\nerror: ", err)
	})

	// 3. called after response received
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited link: ", r.Request.URL)
	})

	// 4. called right after OnResponse if the received content is HTML
	// on every a element which has href attribute call callback
	c.OnHTML("a", func(e *colly.HTMLElement) {
		// printing all URLs associated with the a links in the page
		fmt.Printf("%v", e.Attr("href"))
	})

	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })

	// 5. called after OnHTML callback
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scraping: ", r.Request.URL)
	})

	// downloading the target HTML page
	c.Visit("https://www.scrapingcourse.com/ecommerce/")

}
