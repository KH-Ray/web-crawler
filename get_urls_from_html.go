package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
    hrefs := []string{}

    doc, err := html.Parse(strings.NewReader(htmlBody))
    if err != nil {
        return []string{}, err
    }

    var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					url, err := url.Parse(a.Val)
					if err != nil {
						fmt.Printf("couldn't parse href '%v': %v\n", a.Val, err)
						continue
					}

					if !url.IsAbs() {
						hrefs = append(hrefs, baseURL.ResolveReference(url).String())
					} else {
						hrefs = append(hrefs, url.String())
					}

					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	return hrefs, nil
}
