package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
    argsWithoutProg := os.Args[1:]

    if len(argsWithoutProg) < 1 {
        fmt.Println("no website provided")
        os.Exit(1)
    } else if len(argsWithoutProg) > 3 {
        fmt.Println("too many arguments provided")
        os.Exit(1)
    } 

    rawBaseURL := argsWithoutProg[0]

	maxConcurrency, err := strconv.Atoi(argsWithoutProg[1])
    if err != nil {
        fmt.Println(fmt.Errorf("invalid type for max concurrency -> %v", err))
        return
    }

    maxPages, err := strconv.Atoi(argsWithoutProg[2])
    if err != nil {
        fmt.Println(fmt.Errorf("invalid type for max pages -> %v", err))
        return
    }

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

    printReport(cfg.pages, rawBaseURL)
}
