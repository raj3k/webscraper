package main

import (
	"fmt"
	"os"
	"webscraper/internal/client"
	"webscraper/internal/counter"
	"webscraper/internal/parse"
)

func main() {
	urls := []string{
		"https://toscrape.com/",
		"https://www.scrapethissite.com/pages/",
		"https://example.com/",
	}
	response, err := client.Get(urls[0])
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	doc := parse.HTMLParse(response)
	text := doc.FullText()

	//text := tokenizer.ParseHTML(response)
	//fmt.Println(text)

	mostFrequent := counter.MostFrequentWords(text, 10)

	fmt.Println("Top 3 most frequent words:")
	for _, wc := range mostFrequent {
		fmt.Printf("%s: %d\n", wc.Word, wc.Count)
	}
}
