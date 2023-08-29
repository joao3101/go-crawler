package main

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCraw(t *testing.T) {
	// Create a test HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../internal/test/index.html")
	})

	port := 8080
	fmt.Printf("Server started on port %d\n", port)
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			panic(err)
		}
	}()

	url, err := url.Parse("http://localhost:8080/")
	if err != nil {
		panic(err)
	}

	crawler := NewCrawler()
	crawledURLs := crawler.crawl(url)
	assert.Equal(t, 3, crawledURLs)
}
