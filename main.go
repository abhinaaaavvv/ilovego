package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	Id     int
	Status string
	mu     sync.Mutex
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	orderChan := make(chan *Order, 20)

	go func() {
		defer wg.Done()

		for _, order := range generateOrders(20) {
			orderChan <- order
		}

		close(orderChan)

		fmt.Println("Done with generating orders.")
	}()

	go processOrders(orderChan, &wg)

	wg.Wait()

	fmt.Println("All operations completed.")
}

func processOrders(orderChan <-chan *Order, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orderChan {
		time.Sleep(
			time.Duration(rand.Intn(500)) * time.Millisecond,
		)
		fmt.Printf("Processing order %v\n", order.Id)
	}
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)

	for i := range count {
		orders[i] = &Order{
			Id:     i + 1,
			Status: "Pending",
		}
	}

	return orders
}
