package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	"webscraper/internal/client"
	"webscraper/internal/counter"
	"webscraper/internal/errors"
	"webscraper/internal/parse"
)

type ScrapeResult struct {
	URL    string
	Result string
}

func main() {
	start := time.Now()

	urlsEnv := os.Getenv("URLS")

	urls, err := parseURLEnv(urlsEnv)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	var wg sync.WaitGroup
	resultChan := make(chan ScrapeResult, len(urls))
	for _, url := range urls {
		wg.Add(1)
		go ScrapeUrl(url, &wg, resultChan)
	}

	wg.Wait()

	close(resultChan)

	for sr := range resultChan {
		mostFrequent := counter.MostFrequentWords(sr.Result, 3)
		fmt.Printf("Top 3 most frequent words from %s:\n", sr.URL)
		for _, wc := range mostFrequent {
			fmt.Printf("%s: %d\n", wc.Word, wc.Count)
		}
	}

	// For cache testing
	url := "https://example.com/"
	response, err := client.Get(url)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	doc := parse.HTMLParse(response)
	text := doc.FullText()

	mostFrequent := counter.MostFrequentWords(text, 3)
	fmt.Printf("Top 3 most frequent words from %s:\n", url)
	for _, wc := range mostFrequent {
		fmt.Printf("%s: %d\n", wc.Word, wc.Count)
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

func parseURLEnv(env string) ([]string, error) {
	if env == "" {
		// default websites to scrape if not specified by user
		urls := []string{
			"https://toscrape.com/",
			"https://www.scrapethissite.com/pages/",
			"https://example.com/",
		}
		return urls, nil
	}

	urls := strings.Split(env, ",")

	for i, url := range urls {
		urls[i] = strings.TrimSpace(url)
	}

	for _, url := range urls {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			return nil, errors.NewError(errors.ErrUnableToParse, "Invalid URL format")
		}
	}

	return urls, nil
}
