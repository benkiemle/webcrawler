package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLSFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	node, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	tags := make([]*html.Node, 0)
	findAnchorTags(node, &tags)

	urls := make([]string, len(tags))
	for i, tag := range tags {
		for _, attr := range tag.Attr {
			if attr.Key == "href" {
				test, err := url.Parse(attr.Val)
				if err != nil {
					fmt.Println("an error occurred:", err)
					continue
				}
				if len(test.Host) == 0 {
					urls[i] = rawBaseURL + attr.Val
				} else {
					urls[i] = attr.Val
				}

			}
		}
	}
	return urls, nil
}

func findAnchorTags(node *html.Node, anchorTags *[]*html.Node) {
	if node.Type == html.ElementNode && node.DataAtom == atom.A {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				*anchorTags = append(*anchorTags, node)
			}
		}
	}

	for child := range node.ChildNodes() {
		findAnchorTags(child, anchorTags)
	}
}
