package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	maxConcurrency := 3
	maxPages := 10

	var err error

	if len(os.Args) >= 3 {
		maxConcurrency, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid maxConcurrency value entered:", err)
		}
	}

	if len(os.Args) == 4 {
		maxPages, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("invalid maxPages value entered:", err)
		}
	}

	cfg, err := configure(os.Args[1], maxConcurrency, maxPages)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s...\n", cfg.baseURL.String())

	cfg.wg.Add(1)
	go cfg.crawlPage(cfg.baseURL.String())
	cfg.wg.Wait()

	// for normalizedURL, count := range cfg.pages {
	// 	fmt.Printf("%d - %s\n", count, normalizedURL)
	// }

	cfg.printReport()
}
