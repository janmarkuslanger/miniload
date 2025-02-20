# Mini Load - Lightweight HTTP Load Testing CLI

Mini Load is a simple and lightweight command-line tool for performing HTTP load tests. It allows you to send multiple concurrent requests to a specified URL and measures response times, success rates, and failures.

## Features
- Send concurrent HTTP requests
- Measure response times and calculate averages
- Track success and failure rates
- Configurable concurrency and total request count

## Installation

You can install Mini Load using `go install`:

```sh
go install github.com/yourusername/miniload@latest
```

Ensure that `$GOBIN` is in your `PATH` so you can run `miniload` from anywhere:

```sh
export PATH=$PATH:$(go env GOPATH)/bin
```

## Usage

Run Mini Load with the following command:

```sh
miniload -url "https://example.com" -c 10 -t 100
```

### Parameters:
- `-url` (**required**): The target URL to test.
- `-c` (**optional**, default: 1): Number of concurrent requests.
- `-t` (**optional**, default: 10): Total number of requests to send.

### Example:
```sh
miniload -url "https://api.example.com" -c 5 -t 50
```
This will send 50 requests to `https://api.example.com`, with a concurrency level of 5.

## Output Example
```sh
--- miniload ---
Total requests: 100
Requests succeeded: 95
Requests failed: 5
Average duration: 120ms
--- END ---
```

## License
This project is licensed under the MIT License.

## Contributions
Contributions are welcome! Feel free to submit issues or pull requests to improve Mini Load.

