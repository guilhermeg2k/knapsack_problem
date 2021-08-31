package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGreedy1Solver(t *testing.T) {
	items := readItemsFromFile("example_items/1.data")

	startTime := time.Now()
	knapsack, totalValue, totalWeight := greedyKnapsackSolve(items, 100)
	elapsedTime := time.Since(startTime)

	fmt.Println(elapsedTime)
	PrintMemUsage()

	printItems(knapsack)
	fmt.Println(totalValue, totalWeight)
}
