package main

import "fmt"

func greedyKnapsackSolve(items []Item, knapSackCapacity int) ([]Item, int, int) {
	var knapsack []Item
	totalWeight := 0
	totalValue := 0

	for i := 0; i < len(items); i++ {
		items[i].ValueByWeight = float64(items[i].Value) / float64(items[i].Weight)
	}

	mergeSort(items)

	for i := 0; i < len(items); i++ {
		if (totalWeight + items[i].Weight) < knapSackCapacity {
			knapsack = append(knapsack, items[i])
			totalValue += items[i].Value
			totalWeight += items[i].Weight
		}
		if totalWeight == knapSackCapacity {
			break
		}
	}

	return knapsack, totalValue, totalWeight
}

func main() {
	items := readItemsFromFile("example_items/4.data")
	knapsack, totalValue, totalWeight := greedyKnapsackSolve(items, 10000)
	printItems(knapsack)
	fmt.Println(totalValue, totalWeight)
}
