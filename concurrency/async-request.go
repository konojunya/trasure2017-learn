package main

import (
	"log"
	"net/http"
)

func async() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org",
		"http://www.google.com",
		"http://example.com",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			http.Get(url)
			log.Println(url)
		}(url)
	}
	wg.Wait()
}

func sync() {
	var urls = []string{
		"http://www.golang.org",
		"http://www.google.com",
		"http://example.com",
	}
	for _, url := range urls {
		http.Get(url)
		log.Println(url)
	}
}
