package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("error occurred retrieving url, got status code: %d, %s", resp.StatusCode, resp.Status)
	}

	contentType := resp.Header.Get("content-type")
	if !strings.HasPrefix(contentType, "text/html") {
		return "", fmt.Errorf("invalid content-type found: %s", contentType)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
