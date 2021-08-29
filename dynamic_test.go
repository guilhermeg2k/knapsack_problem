package main

import (
	"fmt"
	"testing"
)

func TestDynamicSolver(t *testing.T) {
	items := readItemsFromFile("example_items/4.data")
	knapsack, totalValue, totalWeight := dynamicKnapsackSolve(items, 100000)
	printItems(knapsack)
	fmt.Println(totalValue, totalWeight)
}
