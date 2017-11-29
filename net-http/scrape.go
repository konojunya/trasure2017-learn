package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type Page struct {
	Title       string
	Description string
	Url         string
	OgTitle     string
	OgImage     string
}

func GetTitleAndDescription(urls []string) []*Page {

	var pages []*Page

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			log.Println(url)
			page, err := Get(url)
			if err != nil {
				fmt.Println(err)
			}
			pages = append(pages, page)
			wg.Done()
		}(url)
	}
	wg.Wait()

	return pages
}

func Get(url string) (*Page, error) {

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	p := Page{
		Title:       getTitle(doc),
		Description: getDescription(doc),
		OgTitle:     getOgTitle(doc),
		OgImage:     getOgImage(doc),
		Url:         url,
	}

	return &p, nil
}

func getTitle(doc *html.Node) string {
	var f func(n *html.Node) string
	f = func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "title" {
			return n.FirstChild.Data
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if s := f(c); s != "" {
				return s
			}
		}

		return ""
	}

	return f(doc)
}

func getDescription(doc *html.Node) string {
	var f func(*html.Node) string
	f = func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "meta" && isDescription(n.Attr) {
			for _, attr := range n.Attr {
				if attr.Key == "content" {
					return attr.Val
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if s := f(c); s != "" {
				return s
			}
		}

		return ""
	}
	return f(doc)
}

func getOgTitle(doc *html.Node) string {
	var f func(*html.Node) string
	f = func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "meta" && isOgTitle(n.Attr) {
			for _, attr := range n.Attr {
				if attr.Key == "content" {
					return attr.Val
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if s := f(c); s != "" {
				return s
			}
		}

		return ""

	}

	return f(doc)
}

func getOgImage(doc *html.Node) string {
	var f func(*html.Node) string
	f = func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "meta" && isOgImage(n.Attr) {
			for _, attr := range n.Attr {
				if attr.Key == "content" {
					return attr.Val
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if s := f(c); s != "" {
				return s
			}
		}

		return ""

	}

	return f(doc)
}

func isDescription(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "name" && attr.Val == "description" {
			return true
		}
	}
	return false
}

func isOgTitle(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "property" && attr.Val == "og:title" {
			return true
		}
	}
	return false
}

func isOgImage(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "property" && attr.Val == "og:image" {
			return true
		}
	}
	return false
}
