package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	cfg.mu.Lock()
	if len(cfg.pages) > cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	currentUrl, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	if cfg.baseURL.Host != currentUrl.Host {
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	isFirst := cfg.addPageVisit(normalizedCurrentURL)
	if !isFirst {
		return
	}

	htmlContent, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(htmlContent)
	fmt.Println(cfg.baseURL.String())
	urls, err := getURLSFromHTML(htmlContent, cfg.baseURL.String())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, u := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(u)
	}
}

func (cfg *config) addPageVisit(noralizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, ok := cfg.pages[noralizedURL]; !ok {
		cfg.pages[noralizedURL] = 1
		isFirst = true
	} else {
		cfg.pages[noralizedURL]++
		isFirst = false
	}
	return isFirst
}
