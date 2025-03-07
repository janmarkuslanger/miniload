package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type RequestResult struct {
	status   int
	duration time.Duration
	error    bool
}

func sendRequest(url string, method string, header map[string]string, client *http.Client) RequestResult {
	start := time.Now()
	request, err := http.NewRequest(method, url, nil)

	print(method)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return RequestResult{error: true}
	}

	for key, value := range header {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return RequestResult{error: true}
	}

	duration := time.Since(start)

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

func parseHeader(headerString string) map[string]string {
	header := make(map[string]string)

	if headerString == "" {
		return header
	}

	pairs := strings.Split(headerString, ";")

	for _, pair := range pairs {
		parts := strings.Split(pair, ":")

		if len(parts) != 2 {
			continue
		}

		header[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	return header
}

func main() {
	url := flag.String("url", "", "URL")
	concurrent := flag.Int("c", 1, "Concurrent requests")
	total := flag.Int("t", 10, "Total requests")
	method := flag.String("m", "GET", "Method can be GET,POST,PUT,...")
	headerFlags := flag.String("h", "", "Header via format: 'Key:Value;Key2:Value2'")

	flag.Parse()

	fmt.Println("--- miniload ---")

	client := &http.Client{Timeout: 5 * time.Second}

	var wg sync.WaitGroup
	var mu sync.Mutex

	results := make([]RequestResult, 0, *total)

	parsedHeader := parseHeader(*headerFlags)

	for i := 0; i < *total; i++ {
		if i%*concurrent == 0 {
			time.Sleep(100 * time.Millisecond)
		}

		wg.Add(1)

		go func() {
			defer wg.Done()
			result := sendRequest(*url, *method, parsedHeader, client)

			mu.Lock()
			results = append(results, result)
			mu.Unlock()
		}()
	}

	wg.Wait()
	printResult(results)
	fmt.Println("--- END ---")
}
