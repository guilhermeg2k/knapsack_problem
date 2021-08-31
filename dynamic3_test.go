package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDynamic3Solver(t *testing.T) {
	items := readItemsFromFile("example_items/3.data")

	startTime := time.Now()
	knapsack, totalValue, totalWeight := dynamicKnapsackSolve(items, 10000)
	elapsedTime := time.Since(startTime)

	fmt.Println(elapsedTime)
	PrintMemUsage()

	printItems(knapsack)
	fmt.Println(totalValue, totalWeight)
}
