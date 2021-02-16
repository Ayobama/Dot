package scraper

import (
	"github.com/gocolly/colly"
)

// The main entity. A collector which manages the nettwork communication and is responsible for the execution of the attached callbacks while a collector job is running.
var c = colly.NewCollector()
fmt.Println(c)
