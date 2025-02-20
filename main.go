package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type RequestResult struct {
	status   int
	duration time.Duration
	error    bool
}

func sendRequest(url string, client *http.Client) RequestResult {
	start := time.Now()
	response, err := client.Get(url)
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return RequestResult{error: true}
	}

	defer response.Body.Close()

	return RequestResult{status: response.StatusCode, duration: duration}
}

func printResult(results []RequestResult) {
	totalCount := len(results)
	var totalDuration time.Duration
	successCount := 0

	for _, item := range results {
		if !item.error {
			totalDuration += item.duration
			successCount++
		}
	}

	averageDuration := totalDuration / time.Duration(successCount)
	fmt.Printf("Total requests: %d\n", totalCount)
	fmt.Printf("Requests succeed: %d\n", successCount)
	fmt.Printf("Requests failed: %d\n", totalCount-successCount)
	fmt.Printf("Average duration: %v\n", averageDuration)
}

func main() {
	url := flag.String("url", "", "URL")
	concurrent := flag.Int("c", 1, "Concurrent requests")
	total := flag.Int("t", 10, "Total requests")

	flag.Parse()

	fmt.Println("--- miniload ---")

	client := &http.Client{Timeout: 5 * time.Second}

	var wg sync.WaitGroup
	var mu sync.Mutex

	results := make([]RequestResult, 0, *total)

	for i := 0; i < *total; i++ {
		if i%*concurrent == 0 {
			time.Sleep(100 * time.Millisecond)
		}

		wg.Add(1)

		go func() {
			defer wg.Done()
			result := sendRequest(*url, client)

			mu.Lock()
			results = append(results, result)
			mu.Unlock()
		}()
	}

	wg.Wait()
	printResult(results)
	fmt.Println("--- END ---")
}
