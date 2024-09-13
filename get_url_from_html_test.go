package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
    tests := []struct{
        name string
        inputURL string
        inputBody string
        expected []string
    }{
        {
            name:     "absolute and relative URLs",
            inputURL: "https://blog.boot.dev",
            inputBody: `
        <html>
            <body>
                <a href="/path/one">
                    <span>Boot.dev</span>
                </a>
                <a href="https://other.com/path/one">
                    <span>Boot.dev</span>
                </a>
            </body>
        </html>
        `,
            expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
        },
        {
            name:     "invalid URLs",
            inputURL: "https://blog.boot.dev",
            inputBody: `
        <html>
            <body>
                <a href=":\\invalidURL">
                    <span>Boot.dev</span>
                </a>
                <a href="#">
                    <span>Empty URL</span>
                </a>
                <a href="https://other.com/path/one">
                    <span>Boot.dev</span>
                </a>
            </body>
        </html>
        `,
            expected: []string{"https://blog.boot.dev","https://other.com/path/one"},
        },
        {
            name:     "no anchor tags",
            inputURL: "https://blog.boot.dev",
            inputBody: `
        <html>
            <body>
                <span>Boot.dev</span>
                <span>Boot.dev</span>
            </body>
        </html>
        `,
            expected: []string{},
        },
        {
            name:     "explicit relative & parent directory URLs",
            inputURL: "https://blog.boot.dev",
            inputBody: `
        <html>
            <body>
                <a href="./path/one">
                    <span>Boot.dev</span>
                </a>
                <a href="../path/two">
                    <span>Boot.dev</span>
                </a>
            </body>
        </html>
        `,
            expected: []string{"https://blog.boot.dev/path/one", "https://blog.boot.dev/path/two"},
        },
    }

    for i, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}
            
            actual, err := getURLsFromHTML(tc.inputBody, baseURL)
            if err != nil {
                t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
                return
            }
            if !reflect.DeepEqual(actual, tc.expected) {
                t.Errorf("Test %v - '%s' FAIL: expected %v, got %v", i, tc.name, tc.expected, actual)
            }
        })
    }
}