package main

import (
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
	resultContainer map[string]interface{}
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

	if data, _ := w.readCache(req.URL.String()); data != "" {
		w.saveToContainer(req.URL.String(), data)
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

	w.updateCache(req.URL.String(), body)

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

type Cached struct {
	Data []byte
}

func (w *WebScraper) readCache(url string) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	cacheDir := filepath.Join(path, "cache")
	sum := sha1.Sum([]byte(url))
	hash := hex.EncodeToString(sum[:])
	if _, err := os.Stat(cacheDir); err != nil {
		return "", err
	}
	cacheFile := filepath.Join(cacheDir, hash)

	file, err := os.Open(cacheFile)
	if err != nil {
		return "", err
	}

	cached := new(Cached)
	err = gob.NewDecoder(file).Decode(cached)
	file.Close()
	return string(cached.Data), nil
}

func (w *WebScraper) updateCache(url string, data []byte) error {
	cached := Cached{
		Data: data,
	}

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	cacheDir := filepath.Join(path, "cache")
	sum := sha1.Sum([]byte(url))
	hash := hex.EncodeToString(sum[:])
	if _, err := os.Stat(cacheDir); err != nil {
		if err := os.MkdirAll(cacheDir, 0750); err != nil {
			return err
		}
	}
	cacheFile := filepath.Join(cacheDir, hash)

	file, err := os.Create(cacheFile)
	if err != nil {
		return err
	}
	if err := gob.NewEncoder(file).Encode(cached); err != nil {
		file.Close()
		return err
	}
	file.Close()

	return nil
}
