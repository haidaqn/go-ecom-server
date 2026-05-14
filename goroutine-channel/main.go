package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	fmt.Println("Starting......")
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	start := time.Now()

	var wg sync.WaitGroup

	for _, id := range ids {
		wg.Add(1)
		go getProductByIdAPI(id, &wg)
	}
	wg.Wait()
	fmt.Println("Done in:", time.Since(start))

}

func getProductByIdAPI(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("https://fakestoreapi.com/products/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(":>>>>>>> Data product by id:", id, "is:", string(body))
}
