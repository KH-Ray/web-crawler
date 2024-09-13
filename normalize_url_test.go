package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct{
		name          string
		inputURL      string
		expected      string
        expectError bool
	}{
		{
			name:     "remove scheme HTTPS",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
            expectError: false,
		},
        {
            name: "remove scheme HTTP",
            inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
            expectError: false,
        },
        {
            name: "handle trailing slash",
            inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
            expectError: false,
        },
        {
            name: "handle uppercase URL",
            inputURL: "HTTPS://BLOG.BOOT.DEV/PATH",
			expected: "blog.boot.dev/path",
            expectError: false,
        },
        {
            name: "handle URL with port number",
            inputURL: "https://blog.boot.dev:80/path",
			expected: "blog.boot.dev/path",
            expectError: false,
        },
        {
            name:     "handle invalid URL",
            inputURL: "not a url",
            expected: "",
            expectError: true,
        },
	}

    for i, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            actual, err := normalizeURL(tc.inputURL)
            if tc.expectError {
                if err == nil {
                    t.Errorf("Test %v - '%s' FAIL: expected an error, but got none", i, tc.name)
                }
            } else {
                if err != nil {
                    t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
                    return
                }
                if actual != tc.expected {
                    t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
                }
            }
        })
    }
}