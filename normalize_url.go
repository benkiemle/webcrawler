package main

import (
	"net/url"
	"strings"
)

func normalizeURL(val string) (string, error) {
	parsedUrl, err := url.Parse(val)
	if err != nil {
		return "", err
	}
	return parsedUrl.Host + strings.TrimRight(parsedUrl.Path, "/"), nil
}
