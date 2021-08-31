package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGreedy2Solver(t *testing.T) {
	items := readItemsFromFile("example_items/2.data")

	startTime := time.Now()
	knapsack, totalValue, totalWeight := greedyKnapsackSolve(items, 1000)
	elapsedTime := time.Since(startTime)

	fmt.Println(elapsedTime)
	PrintMemUsage()

	printItems(knapsack)
	fmt.Println(totalValue, totalWeight)
}
