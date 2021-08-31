package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDynamic4Solver(t *testing.T) {
	items := readItemsFromFile("example_items/4.data")

	startTime := time.Now()
	knapsack, totalValue, totalWeight := dynamicKnapsackSolve(items, 100000)
	elapsedTime := time.Since(startTime)

	fmt.Println(elapsedTime)
	PrintMemUsage()

	printItems(knapsack)
	fmt.Println(totalValue, totalWeight)
}
