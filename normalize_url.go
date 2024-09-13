package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(urlString string) (string, error) {
    parsedURL, err := url.Parse(urlString)
    if err != nil {
        return "", err
    }

    normalizedURL := strings.ToLower(parsedURL.Host)

    if parsedURL.Host == "" {
        return "", errors.New("invalid URL")
    }

    if parsedURL.Path != "" {
        normalizedURL += strings.TrimSuffix(parsedURL.Path, "/")
    } 

    if strings.Contains(parsedURL.Host, parsedURL.Port()) {
        normalizedURL = strings.Replace(normalizedURL, fmt.Sprintf(":%v", parsedURL.Port()), "", 1)
    }

    normalizedURL = strings.Replace(normalizedURL, parsedURL.Path, strings.ToLower(parsedURL.Path), 1)
    
    return normalizedURL, nil
}
