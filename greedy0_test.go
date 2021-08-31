package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGreedy0Solver(t *testing.T) {
	items := readItemsFromFile("example_items/0.data")

	startTime := time.Now()
	knapsack, totalValue, totalWeight := greedyKnapsackSolve(items, 10)
	elapsedTime := time.Since(startTime)

	fmt.Println(elapsedTime)
	PrintMemUsage()

	printItems(knapsack)
	fmt.Println(totalValue, totalWeight)
}
