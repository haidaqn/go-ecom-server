package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	urls := []string{
		"http://127.0.0.1:8089/action/stock?msg=HSA&type=MUA",
		"http://127.0.0.1:8089/action/stock?msg=HSA&type=BAN",
	}

	start := time.Now()

	// gửi 10 request cùng lúc
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			url := urls[index%len(urls)]

			resp, err := http.Post(url, "application/json", nil)
			if err != nil {
				fmt.Printf("Request %d lỗi: %v\n", index, err)
				return
			}
			defer resp.Body.Close()

			body, _ := io.ReadAll(resp.Body)

			fmt.Printf(
				"Request %d | Status: %d | Body: %s\n",
				index,
				resp.StatusCode,
				string(body),
			)
		}(i)
	}

	wg.Wait()

	fmt.Println("Done in:", time.Since(start))
}
