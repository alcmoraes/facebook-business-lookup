package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

var duckGoSearchEndpoint = "https://duckduckgo.com/html/?q=%s"

func main() {
	firstMatch := 0
	c := colly.NewCollector()

	c.OnHTML("table div", func(e *colly.HTMLElement) {
		if firstMatch == 1 {
			if strings.Contains(e.Text, "Ligar ") {
				fmt.Println(e.Text)
				os.Exit(0)
			}
		}
	})

	c.OnHTML(".result__body", func(e *colly.HTMLElement) {
		if firstMatch == 0 {
			firstMatch = 1
			fmt.Println(e.ChildText(".result__a"))
			fmt.Println("https://m." + e.ChildText(".result__url"))
			e.Request.Visit("https://m." + e.ChildText(".result__url"))
		}
		firstMatch = -1
	})
	c.Visit(fmt.Sprintf(duckGoSearchEndpoint, os.Args[1]+" Facebook"))

}
