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
	wg.Add(3)

	orderChan := make(chan *Order, 20)
	processedChan := make(chan *Order, 20)

	go func() {
		defer wg.Done()
		for _, order := range generateOrders(20) {
			orderChan <- order
		}
		close(orderChan)
		fmt.Println("Done with generating orders.")
	}()

	go processOrders(orderChan, processedChan, &wg)

	go func() {
		defer wg.Done()
		for {
			select {
			case processOrder, ok := <-processedChan:
				if !ok {
					fmt.Println("Processing channel closed")
					return
				}
				fmt.Printf("Processed order %v with status: %s\n", processOrder.Id, processOrder.Status)

			case <-time.After(10 * time.Second):
				fmt.Println("Timeout waiting for operations.")
				return
			}
		}
	}()

	wg.Wait()
	fmt.Println("All operations completed.")
}

func processOrders(inChan <-chan *Order, outChan chan<- *Order, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		close(outChan)
	}()

	for order := range inChan {
		time.Sleep(
			time.Duration(rand.Intn(500)) * time.Millisecond,
		)
		order.Status = "Processed"
		outChan <- order
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
