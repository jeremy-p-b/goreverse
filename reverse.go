// Package reverse provides a simple reverse look-up for URLS.
package goreverse

import "fmt"

type URLReverse struct {
	urls         map[string]string
	ajaxPostURLS []string
}

// NewURLReverse creates a new empty URLReverse
func NewURLReverse() URLReverse {
	return URLReverse{
		urls: make(map[string]string),
	}
}

// Add adds a named url to the URLReverse
// Panics if url already added to URLReverse
func (u *URLReverse) Add(urlName, url string) string {
	if u.urls[urlName] != "" && u.urls[urlName] != url {
		panicMessage := fmt.Sprintf("Panic: the url %s has already been added to URLReverse", url)
		panic(panicMessage)
	}
	u.urls[urlName] = url
	return url
}

// Get gets a named url from the URLReverse
func (u *URLReverse) Get(urlName string) string {
	return u.urls[urlName]
}
