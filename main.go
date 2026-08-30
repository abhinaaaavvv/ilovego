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

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	orders := generateOrders(20)

	// go func()  {
	// 		processOrders(orders)
	// 		defer wg.Done()
	// }()

	for range 3 {
		go func() {
			for _, order := range orders {
				updateOrderStatuses(order)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	reportOrderStatus(orders)
	fmt.Println("All operations completed.")
	fmt.Println(totalUpdates)
}

func updateOrderStatuses(order *Order) {
	order.mu.Lock()
	time.Sleep(
		time.Duration(rand.Intn(300)) * time.Millisecond,
	)

	status := []string{
		"Pending", "Shipped", "Delivered",
	}[rand.Intn(3)]

	order.Status = status

	fmt.Printf(
		"Updated order %v status: %s\n",
		order.Id, status,
	)
	order.mu.Unlock()

	updateMutex.Lock()
	defer updateMutex.Unlock()
	currentUpdates := totalUpdates
	time.Sleep(5 * time.Millisecond)
	totalUpdates = currentUpdates + 1
}

func processOrders(orders []*Order) {
	for _, order := range orders {
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

func reportOrderStatus(orders []*Order) {
	fmt.Println("\n--- Order Status Report ---")

	for _, order := range orders {
		fmt.Printf(
			"Order %v: %s\n",
			order.Id, order.Status,
		)
	}
	fmt.Println("---------------------------")
}
