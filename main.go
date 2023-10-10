package main

import (
	"fmt"
	"sync"
	"time"
	"webscraper/internal/client"
	"webscraper/internal/counter"
	"webscraper/internal/parse"
)

type ScrapeResult struct {
	URL    string
	Result string
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	urls := []string{
		"https://toscrape.com/",
		"https://www.scrapethissite.com/pages/",
		"https://example.com/",
	}

	resultChan := make(chan ScrapeResult, len(urls))
	for _, url := range urls {
		wg.Add(1)
		go ScrapeUrl(url, &wg, resultChan)
	}

	wg.Wait()

	close(resultChan)

	for sr := range resultChan {
		//mostFrequent = append(mostFrequent, counter.MostFrequentWords(result, 3))
		mostFrequent := counter.MostFrequentWords(sr.Result, 3)
		fmt.Printf("Top 3 most frequent words from %s:\n", sr.URL)
		for _, wc := range mostFrequent {
			fmt.Printf("%s: %d\n", wc.Word, wc.Count)
		}
	}

	fmt.Println(time.Since(start))
}

func ScrapeUrl(url string, wg *sync.WaitGroup, resultChan chan<- ScrapeResult) {
	defer wg.Done()

	response, err := client.Get(url)
	if err != nil {
		resultChan <- ScrapeResult{URL: url, Result: fmt.Sprintf("Error scraping %s: %s", url, err.Error())}
	}

	doc := parse.HTMLParse(response)
	text := doc.FullText()
	//text := tokenizer.ParseHTML(response)

	resultChan <- ScrapeResult{URL: url, Result: text}
}
