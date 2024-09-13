package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
    cfg.concurrencyControl <- struct{}{}
    defer func() {
        <-cfg.concurrencyControl
        cfg.wg.Done()
    }()

    if validLength := cfg.checkLengthOfPages(); !validLength {
        return
    }

    parsedCurrentURL, err := url.Parse(rawCurrentURL)
    if err != nil {
        fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
    }

    if parsedCurrentURL.Host != cfg.baseURL.Host {
        return
    }

    normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
		return
	}

    if isFirst := cfg.addPageVisit(normalizedURL); !isFirst {
        return
    }

    fmt.Printf("crawling %s\n", rawCurrentURL)

    htmlBody, err := getHTML(rawCurrentURL)
    if err != nil {
        fmt.Printf("Error - getHTML: %v", err)
		return
    }

    urls, err := getURLsFromHTML(htmlBody, cfg.baseURL)
    if err != nil {
        fmt.Printf("Error - getURLsFromHTML: %v", err)
		return
    }

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
