package scraping

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)


func DetermineUrls() {
	collyScraper := colly.NewCollector()


	collyScraper.OnHTML("main", func(e *colly.HTMLElement) {
		println("Printing out main!")
		
	})
	collyScraper.OnRequest(func(r *colly.Request){
		fmt.Println("Visting", r.URL.String())
	})

	collyScraper.OnScraped(func(r *colly.Response){
		fmt.Println(e)
		
	})
	collyScraper.Visit("https://en.wikipedia.org/wiki/Jiaozi")


}