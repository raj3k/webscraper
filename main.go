package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
	"webscraper/internal/counter"
	"webscraper/internal/errors"
	"webscraper/internal/parse"
)

func main() {
	start := time.Now()
	urlsEnv := os.Getenv("URLS")
	CPUCount := runtime.NumCPU()

	ws := NewWebScraper(Async())

	ws.LimitGoroutines(CPUCount)

	urls, err := parseURLEnv(urlsEnv)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	for _, url := range urls {
		ws.Scrape(url)
	}

	ws.Wait()

	for url, data := range ws.resultContainer {
		doc := parse.HTMLParse(data.(string))
		text := doc.FullText()
		mostFrequent := counter.MostFrequentWords(text, 5)
		fmt.Printf("Top %d most frequent words from %s:\n", 5, url)
		for _, wc := range mostFrequent {
			fmt.Printf("%s: %d\n", wc.Word, wc.Count)
		}
	}

	urls = append(urls, "https://toscrape.com/", "https://www.scrapethissite.com/pages/")

	for _, url := range urls {
		ws.Scrape(url)
	}

	ws.Wait()

	for url, data := range ws.resultContainer {
		doc := parse.HTMLParse(data.(string))
		text := doc.FullText()
		mostFrequent := counter.MostFrequentWords(text, 5)
		fmt.Printf("Top %d most frequent words from %s:\n", 5, url)
		for _, wc := range mostFrequent {
			fmt.Printf("%s: %d\n", wc.Word, wc.Count)
		}
	}

	fmt.Println("------------------------------------------")
	fmt.Println("Running time:", time.Since(start))
}

func parseURLEnv(env string) ([]string, error) {
	if env == "" {
		// default websites to scrape if not specified by user
		urls := []string{
			"https://example.com/",
			"https://www.wp.pl",
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
