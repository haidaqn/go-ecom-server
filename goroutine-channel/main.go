package main

// import (
// 	"fmt"
// 	"time"
// )

// // import (
// // 	"fmt"
// // 	"io"
// // 	"net/http"
// // 	"sync"
// // 	"time"
// // )

// // func main() {
// // 	fmt.Println("Starting......")
// // 	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// // 	start := time.Now()

// // 	var wg sync.WaitGroup

// // 	for _, id := range ids {
// // 		wg.Add(1)
// // 		go getProductByIdAPI(id, &wg)
// // 	}
// // 	wg.Wait()
// // 	fmt.Println("Done in:", time.Since(start))

// // }

// // func getProductByIdAPI(id int, wg *sync.WaitGroup) {
// // 	defer wg.Done()

// // 	url := fmt.Sprintf("https://fakestoreapi.com/products/%d", id)
// // 	resp, err := http.Get(url)
// // 	if err != nil {
// // 		fmt.Println("Error:", err)
// // 		return
// // 	}
// // 	defer resp.Body.Close()
// // 	body, err := io.ReadAll(resp.Body)
// // 	if err != nil {
// // 		fmt.Println("Error:", err)
// // 		return
// // 	}
// // 	fmt.Println(":>>>>>>> Data product by id:", id, "is:", string(body))
// // }

// // import "fmt"

// // type Course struct {
// // 	title string
// // 	price int
// // }

// // func main() {
// // 	// 1. add channel to

// // 	ch := make(chan Course)

// // 	//2. create goroutine

// // 	go func() {
// // 		course := Course{
// // 			title: "Golang",
// // 			price: 100,
// // 		}
// // 		ch <- course // send data to channel
// // 	}()

// // 	c := <-ch // receive data from channel
// // 	fmt.Println(c)

// //  }

// // pub/sub use channel and gouroutines

// type Message struct {
// 	OrderId string
// 	Title   string
// 	Price   int
// }

// // func publisher(ch chan Message, orders []Message) {
// // 	for _, order := range orders {

// // 		fmt.Printf("Publishing order: %s\n", order.OrderId)
// // 		ch <- order
// // 		time.Sleep(1 * time.Second)
// // 	}
// // 	close(ch)
// // }

// // func subscriber(ch <-chan Message, user string) {
// // 	for order := range ch {
// // 		fmt.Printf("User %s received order: %s price %d\n", user, order.OrderId, order.Price)
// // 		time.Sleep(1 * time.Second)
// // 	}

// // }

// // func main() {
// // 	// 1. channel order
// // 	orderChannle := make(chan Message)

// // 	// 2. create orders
// // 	orders := []Message{
// // 		{OrderId: "1", Title: "Golang", Price: 100},
// // 		{OrderId: "2", Title: "Python", Price: 200},
// // 		{OrderId: "3", Title: "Java", Price: 300},
// // 		{OrderId: "4", Title: "JavaScript", Price: 400},
// // 	}

// // 	// 3. send order to pub
// // 	go publisher(orderChannle, orders)
// // 	go subscriber(orderChannle, "Hai Dang")
// // 	time.Sleep(10 * time.Second)
// // 	// wait for all goroutines to finish

// // }

// func buyTicket(channel chan<- Message, orders []Message) {
// 	for _, order := range orders {
// 		time.Sleep(1 * time.Second)
// 		fmt.Printf("Buying order: %s\n", order.OrderId)
// 		channel <- order // send data to channel
// 	}
// 	close(channel) // close channel after sending all orders
// }

// func cancelTicket(channel chan<- string, orderIds []string) {
// 	for _, orderId := range orderIds {
// 		time.Sleep(10 * time.Second)
// 		fmt.Printf("Canceling order: %s\n", orderId)
// 	}
// 	close(channel)
// }

// func handlerOrder(orderChannel <-chan Message, cancelChannel <-chan string) {
// 	for {
// 		select {
// 		case order, orderOk := <-orderChannel:
// 			if !orderOk {
// 				fmt.Println("No more orders to process.")
// 				orderChannel = nil // set to nil to avoid further reads
// 			} else {
// 				fmt.Printf("Processing order: %s price %d\n", order.OrderId, order.Price)
// 			}
// 		case cancel, cancelOk := <-cancelChannel:
// 			if !cancelOk {
// 				fmt.Println("No more cancellations to process.")
// 				cancelChannel = nil // set to nil to avoid further reads
// 			} else {
// 				fmt.Printf("Processing cancellation: %s\n", cancel)
// 			}
// 		}

// 		// exit
// 		if orderChannel == nil && cancelChannel == nil {
// 			fmt.Println("All orders and cancellations processed. Exiting handler.")
// 			break
// 		}

// 	}
// }

// func main() {
// 	buyChannel := make(chan Message)
// 	cancelChannel := make(chan string)

// 	// simulate order ticket

// 	buyOrders := []Message{
// 		{OrderId: "Order-01", Title: "Golang", Price: 100},
// 		{OrderId: "Order-02", Title: "Python", Price: 200},
// 		{OrderId: "Order-03", Title: "Java", Price: 300},
// 		{OrderId: "Order-04", Title: "JavaScript", Price: 400},
// 	}

// 	cancelOrders := []string{"Order-02", "Order-04"}

// 	go buyTicket(buyChannel, buyOrders)
// 	go cancelTicket(cancelChannel, cancelOrders)
// 	go handlerOrder(buyChannel, cancelChannel)

// 	time.Sleep(15 * time.Second)

// }
