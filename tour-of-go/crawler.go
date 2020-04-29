package main

import (
	"fmt"
	"sync"
)

var waitgroup sync.WaitGroup

type Cache struct {
	v   map[string]bool
	mux sync.Mutex
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func (c *Cache) showInCache(url string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()

	if c.v[url] {
		return true
	}
	return false
}

func (c *Cache) undateCache(url string) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.v[url] = true
	return
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *Cache) Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	if c.showInCache(url) {
		return
	} else {
		c.undateCache(url)
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		func() {
			defer waitgroup.Done()
			waitgroup.Add(1)
			c.Crawl(u, depth-1, fetcher)
		}()
	}
	return
}

func main() {
	//waitgroup.Add(11)

	c := Cache{v: make(map[string]bool)}
	c.Crawl("https://golang.org/", 4, fetcher)
	waitgroup.Wait()

}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}