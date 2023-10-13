package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

type WebScraperOption func(*WebScraper)

func Async(a ...bool) WebScraperOption {
	return func(ws *WebScraper) {
		if len(a) > 0 {
			ws.Async = a[0]
		} else {
			ws.Async = true
		}
	}
}

type WebScraper struct {
	Async bool

	wg              sync.WaitGroup
	mu              sync.Mutex
	waitChan        chan int
	resultContainer map[string]interface{} // works like cache
}

func NewWebScraper(options ...WebScraperOption) *WebScraper {
	scraper := new(WebScraper)

	for _, f := range options {
		f(scraper)
	}

	return scraper
}

func (w *WebScraper) Scrape(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if w.Async {
		w.wg.Add(1)
		w.waitChan <- 1
		go func(url string) {
			defer w.wg.Done()
			w.fetch(req)
			<-w.waitChan
		}(url)
		return nil
	}

	return w.fetch(req)
}

func (w *WebScraper) fetch(req *http.Request) error {
	w.mu.Lock()
	_, found := w.resultContainer[req.URL.String()]
	w.mu.Unlock()
	if found {
		fmt.Println("Cache hit for URL:", req.URL.String())
		return nil
	}

	defaultClient := &http.Client{}
	response, err := defaultClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	bodyReader := response.Body

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		return err
	}

	w.saveToContainer(req.URL.String(), string(body))
	return nil
}

func (w *WebScraper) saveToContainer(url string, data string) {
	if w.resultContainer == nil {
		w.resultContainer = make(map[string]interface{})
	}

	if w.Async {
		w.mu.Lock()
		defer w.mu.Unlock()
	}

	w.resultContainer[url] = data
}

func (w *WebScraper) Wait() {
	w.wg.Wait()
}

func (w *WebScraper) LimitGoroutines(limit ...int) {
	var waitChanSize int
	if len(limit) > 0 {
		waitChanSize = limit[0]
	} else {
		waitChanSize = 1
	}
	w.waitChan = make(chan int, waitChanSize)
}
