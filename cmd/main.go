package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/joao3101/go-crawler/internal/config"
	httpClient "github.com/joao3101/go-crawler/internal/http"
)

type crawler struct {
	client      httpClient.HTTPClient
	startingURL *url.URL
	visitedUrls map[string]bool
	mutex       sync.Mutex
	sem         chan struct{}
}

type Crawler interface {
	crawl(url *url.URL)
}

func NewCrawler() *crawler {
	client := http.DefaultClient
	mainURL := config.Config.MainURL
	startingURL, err := url.Parse(mainURL)
	if err != nil {
		panic("failed to parse url")
	}

	maxNumOfGoRoutines := config.Config.MaxNumOfGoRoutines

	return &crawler{
		client:      client,
		startingURL: startingURL,
		visitedUrls: make(map[string]bool),
		sem:         make(chan struct{}, maxNumOfGoRoutines),
	}
}

var _ Crawler = (*crawler)(nil)

func main() {
	crawler := NewCrawler()

	crawler.crawl(crawler.startingURL)
}

func (c *crawler) crawl(url *url.URL) {
	resp, err := c.client.Get(url.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// This returns a doc that can be iterated and this way we can check for href's
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	doc.Find("a").Each(func(index int, element *goquery.Selection) {
		c.sem <- struct{}{}

		wg.Add(1)
		go func(element *goquery.Selection) {
			fmt.Printf("Size of semaphore at the beggining of execution: %d\n", len(c.sem))
			defer func() {
				// Free the semaphore at the end of execution
				<-c.sem
				wg.Done()
				fmt.Printf("Size of semaphore at the end of execution: %d\n", len(c.sem))
			}()
			link, exists := element.Attr("href")
			if exists {
				linkURL, err := url.Parse(link)
				if err != nil {
					return
				}

				if linkURL.Scheme+linkURL.Host == url.Scheme+url.Host {
					c.mutex.Lock()
					if !c.visitedUrls[linkURL.String()] {
						c.visitedUrls[linkURL.String()] = true
						c.mutex.Unlock()

						go c.crawl(linkURL) // Can calling this as a routine give me problems? Check
					} else {
						// Even if there's no match, still need to unlock the mutex
						c.mutex.Unlock()
					}
				}
			}
		}(element)
	})
	// This ensures all goroutines have finished
	wg.Wait()

	// Closes the semaphore channel after we're sure go routines are all done
	close(c.sem)

	for visitedURL := range c.visitedUrls {
		fmt.Println("Visited:", visitedURL)
	}

	fmt.Println("#############################")
	fmt.Println(len(c.visitedUrls))
	fmt.Println("#############################")
}
