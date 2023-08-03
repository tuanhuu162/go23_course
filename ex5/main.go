package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

type NewsArticle struct {
	Title       string
	Topic       string
	Description string
	Imgs        string
	Url         string
}

func main() {
	fName := "data.json"
	file, err := os.OpenFile(fName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Can not create file %q: %s\n", fName, err)
	}
	defer file.Close()

	// init base collector
	c := colly.NewCollector(
		// Visit only domains: https://vnexpress.net/, vnexpress.net
		colly.AllowedDomains("vnexpress.net", "https://vnexpress.net/"),
		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./vnexpress_cache"),
	)

	// init article collector
	articleCollector := c.Clone()

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	// Start fetching the topic on navibar and get list of news in that topic
	c.OnHTML("nav.main-nav li", func(e *colly.HTMLElement) {
		if e.Attr("class") != "home" {
			link := e.ChildAttr("a", "href")
			request := e.Request
			request.Ctx.Put("topic", e.ChildText("a"))
			articleCollector.Visit(request.AbsoluteURL(link))
		}
	})

	articles := make([]NewsArticle, 0, 200)
	articleCollector.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String(), r.Ctx.Get("topic"))
	})
	// Start fetching the new lists in the topic
	articleCollector.OnHTML("article", func(e *colly.HTMLElement) {

		article := NewsArticle{}
		article.Topic = e.Request.Ctx.Get("topic")
		article.Title = e.ChildText("h3.title-news")
		article.Description = e.ChildText("p.description")
		article.Imgs = e.ChildAttr("img", "src")
		article.Url = e.ChildAttr("a", "href")
		if article.Url != "" {
			articles = append(articles, article)
		}
		next_page := e.ChildAttr("a.next-page", "href")
		if next_page != "" {
			e.Request.Visit(e.Request.AbsoluteURL(next_page))
		}
	})

	c.Visit("https://vnexpress.net/")

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(articles)

}
