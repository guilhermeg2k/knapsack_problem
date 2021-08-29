package main

import (
	"fmt"
	"testing"
)

func TestGreedySolver(t *testing.T) {
	items := readItemsFromFile("example_items/4.data")
	knapsack, totalValue, totalWeight := greedyKnapsackSolve(items, 100000)
	printItems(knapsack)
	fmt.Println(totalValue, totalWeight)
}
