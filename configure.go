package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
    maxPages            int
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
    cfg.mu.Lock()
    defer cfg.mu.Unlock()

    if _, exists := cfg.pages[normalizedURL]; !exists {
        cfg.pages[normalizedURL] = 1
        isFirst = true
    } else {
        cfg.pages[normalizedURL]++
        isFirst = false
    }

    return isFirst
}

func (cfg *config) checkLengthOfPages() (validLength bool) {
    cfg.mu.Lock()
    defer cfg.mu.Unlock()

    return len(cfg.pages) < cfg.maxPages
}

func configure(rawBaseURL string, maxConcurrency, maxPages int) (*config, error) {
    baseURL, err := url.Parse(rawBaseURL)
    if err != nil {
        return nil, fmt.Errorf("couldn't parse base URL: %v", err)
    }

    return &config{
        pages: make(map[string]int),
        baseURL: baseURL,
        mu: &sync.Mutex{},
        concurrencyControl: make(chan struct{}, maxConcurrency),
        wg: &sync.WaitGroup{},
        maxPages: maxPages,
    }, nil
}
