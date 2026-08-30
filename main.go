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
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	orders := generateOrders(20)

	go func() {
		processOrders(orders)
		defer wg.Done()
	}()

	go func() {
		updateOrderStatuses(orders)
		defer wg.Done()
	}()

	go func() {
		reportOrderStatus(orders)
		defer wg.Done()
	}()

	wg.Wait()

	fmt.Println("All operations completed.")
}

func updateOrderStatuses(orders []*Order) {
	for _, order := range orders {
		time.Sleep(
			time.Duration(rand.Intn(300)) * time.Millisecond,
		)

		status := []string{
			"Pending", "Shipped", "Delivered",
		}[rand.Intn(3)]

		order.Status = status

		fmt.Printf(
			"Updated order %v status: %s\n",
			order.Id, order.Status,
		)
	}
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
			i + 1, "Pending",
		}
	}

	return orders
}

func reportOrderStatus(orders []*Order) {
	for range 5 {
		time.Sleep(1 * time.Second)
		fmt.Println("\n--- Order Status Report ---")

		for _, order := range orders {
			fmt.Printf(
				"Order %v: %s\n",
				order.Id, order.Status,
			)
		}
	}
	fmt.Println("---------------------------")
}
