package main

type WebScraperOption func(*WebScraper)

type WebScraper struct {
	Async bool
}

func NewWebScraper(options ...WebScraperOption) *WebScraper {
	scraper := new(WebScraper)

	for _, f := range options {
		f(scraper)
	}

	return scraper
}

func Async(a ...bool) WebScraperOption {
	return func(ws *WebScraper) {
		if len(a) > 0 {
			ws.Async = a[0]
		} else {
			ws.Async = true
		}
	}
}
