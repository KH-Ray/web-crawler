package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
    res, err := http.Get(rawURL)
    if err != nil {
        return "", err
    }
    defer res.Body.Close()
    if res.StatusCode > 299 {
		return "", fmt.Errorf("HTTP request failed with status code: %d for URL: %s", res.StatusCode, rawURL)
	}

    body, err := io.ReadAll(res.Body)
    if err != nil {
		return "", err
	}

    contentType := res.Header.Get("Content-Type")
    if !strings.Contains(contentType, "text/html") {
        return "", fmt.Errorf("invalid Content-Type: got %s, want text/html for URL: %s", contentType, rawURL)
    }

    htmlBody := string(body)

    return htmlBody, nil
}