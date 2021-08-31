package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDynamic1Solver(t *testing.T) {
	items := readItemsFromFile("example_items/1.data")

	startTime := time.Now()
	knapsack, totalValue, totalWeight := dynamicKnapsackSolve(items, 100)
	elapsedTime := time.Since(startTime)

	fmt.Println(elapsedTime)
	PrintMemUsage()

	printItems(knapsack)
	fmt.Println(totalValue, totalWeight)
}
