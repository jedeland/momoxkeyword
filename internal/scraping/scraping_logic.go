package scraping

import (
	"fmt"
	"github.com/gocolly/colly"
)


func DetermineUrls() {
	collyScraper := colly.NewCollector()


	collyScraper.OnHTML("main", func(e *colly.HTMLElement) {
		fmt.Printf("Printing out main!")
		bodyOutput := e.ChildText("div.mw-body-content")
		fmt.Println(bodyOutput)
		
	})
	collyScraper.OnRequest(func(r *colly.Request){
		fmt.Println("Visting", r.URL.String())
	})

	collyScraper.Visit("https://en.wikipedia.org/wiki/Jiaozi")


}