package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

// The main entity. A collector which manages the nettwork communication and is responsible for the execution of the attached callbacks while a collector job is running.
var c = colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))

func main() {
	// Adding a callback to a collector
	// if the collector encounters an error log the error
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting: %s", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Printf("Something went wrong: %s", err.Error())
	})
	// c.OnHTML("")
	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("Finished scraping site: %s", r.Request.URL)
	})
}
