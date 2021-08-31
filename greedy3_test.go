package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGreedy3Solver(t *testing.T) {
	items := readItemsFromFile("example_items/3.data")

	startTime := time.Now()
	knapsack, totalValue, totalWeight := greedyKnapsackSolve(items, 10000)
	elapsedTime := time.Since(startTime)

	fmt.Println(elapsedTime)
	PrintMemUsage()

	printItems(knapsack)
	fmt.Println(totalValue, totalWeight)
}
