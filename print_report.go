package main

import (
	"fmt"
	"net/url"
	"sort"
)

type Page struct {
    URL   string
    Count int
}

func printReport(pages map[string]int, baseURL string) {
    fmt.Println("==============================")
    fmt.Printf("REPORT for %s\n", baseURL)
    fmt.Println("==============================")

    var sortedPages []Page
    for url, count := range pages {
        sortedPages = append(sortedPages, Page{URL: url, Count: count})
    }
    sort.Slice(sortedPages, func(i, j int) bool { 
        if sortedPages[i].Count != sortedPages[j].Count {
            return sortedPages[i].Count > sortedPages[j].Count
        }
        return sortedPages[i].URL < sortedPages[j].URL
    })

    parsedBaseURL, err := url.Parse(baseURL)
    if err != nil {
        fmt.Printf("couldn't parse base URL: %v", err)
        return
    }

    for _, page := range sortedPages {
        parsedPageURL, err := url.Parse(page.URL)
        if err != nil {
            fmt.Printf("couldn't parse page URL: %v", err)
            continue
        }
        resolvedURL := parsedBaseURL.ResolveReference(parsedPageURL)

        fmt.Printf("Found %d internal links to %s\n", page.Count, resolvedURL)
    }
}
