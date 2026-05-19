package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	Url    string
	Status string
	Err    error
}

func checkWebsiteUrl(url string, ch chan Result) {
	// Good practice: Add a timeout so a slow site doesn't block forever
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Get(url)
	if err != nil {
		ch <- Result{
			Url:    url,
			Status: "is down",
			Err:    err,
		}
		return
	}
	defer res.Body.Close()

	ch <- Result{
		Url:    url,
		Status: "is up and running",
		Err:    nil,
	}
}

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://shihabdev.me",
	}

	ch := make(chan Result)
	start := time.Now()

	// Spawn 3 goroutines (First pass)
	for _, url := range urls {
		go checkWebsiteUrl(url, ch)
	}

	// Spawn 3 goroutines (Second pass)
	for _, url := range urls {
		go checkWebsiteUrl(url, ch)
	}

	// FIX: Loop 6 times (len(urls) * 2) because 6 goroutines were started
	totalChecks := len(urls) * 2
	for i := 0; i < totalChecks; i++ {
		result := <-ch // Safely reads all 6 results

		if result.Err != nil {
			fmt.Printf("❌ %s %s | Error: %v\n", result.Url, result.Status, result.Err)
			continue
		}

		fmt.Printf("✅ %s %s\n", result.Url, result.Status)
	}

	fmt.Println("Time taken:", time.Since(start))
	fmt.Println("All URLs checked successfully")
}
