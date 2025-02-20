# Mini Load 

Mini Load is a simple and lightweight Go package for performing HTTP load tests. It allows you to send multiple concurrent requests to a specified URL and measure response times, success rates, and failures.

## Features
- Send concurrent HTTP requests
- Measure response times and calculate averages
- Track success and failure rates
- Configurable concurrency and total request count

## Installation

Install Mini Load using Go modules:

```sh
go get github.com/yourusername/miniload
```

## Usage

Import the package and use it in your Go application:

```go
package main

import (
	"fmt"
	"github.com/yourusername/miniload"
)

func main() {
	results := miniload.RunTest("https://example.com", 10, 100)
	fmt.Println(results)
}
```

### Function Parameters:
- `url` (**string**): The target URL to test.
- `concurrency` (**int**): Number of concurrent requests.
- `totalRequests` (**int**): Total number of requests to send.

### Example:
```go
results := miniload.RunTest("https://api.example.com", 5, 50)
fmt.Println(results)
```
This will send 50 requests to `https://api.example.com`, with a concurrency level of 5.

## Output Example
```sh
Total requests: 100
Requests succeeded: 95
Requests failed: 5
Average duration: 120ms
```

## License
This project is licensed under the MIT License.

## Contributions
Contributions are welcome! Feel free to submit issues or pull requests to improve Mini Load.

