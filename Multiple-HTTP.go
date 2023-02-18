package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{"https://www.google.com", "https://www.facebook.com", "https://www.twitter.com"}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	results := make(chan string)

	for _, url := range urls {
		go func(url string) {
			select {
			case <-ctx.Done():
				return
			default:
				resp, err := http.Get(url)
				if err != nil {
					results <- fmt.Sprintf("%s failed: %s", url, err)
					return
				}
				defer resp.Body.Close()
				results <- fmt.Sprintf("%s returned status code %d", url, resp.StatusCode)
			}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-ctx.Done():
			fmt.Println("Cancelled due to timeout")
			return
		}
	}
}
