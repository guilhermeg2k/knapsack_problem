package main

func dynamicKnapsackSolve(items []Item, knapSackCapacity int) ([]Item, int, int) {
	var knapsack []Item
	table := make([][]int, len(items)+1)
	for i := 0; i < len(items)+1; i++ {
		line := make([]int, knapSackCapacity+1)
		table[i] = line

		for j := 0; j < knapSackCapacity+1; j++ {
			if i == 0 || j == 0 {
				continue
			}
			if items[i-1].Weight <= j {
				aux := items[i-1].Value + table[i-1][j-items[i-1].Weight]
				if aux > table[i-1][j] {
					table[i][j] = aux
				} else {
					table[i][j] = table[i-1][j]
				}
			} else {
				table[i][j] = table[i-1][j]
			}
		}

	}

	currentLine := len(items)
	currentColumn := knapSackCapacity
	totalWeight := 0

	for {
		if table[currentLine][currentColumn] != table[currentLine-1][currentColumn] {
			knapsack = append(knapsack, items[currentLine-1])
			totalWeight += items[currentLine-1].Weight
			currentColumn -= items[currentLine-1].Weight
		}
		currentLine--
		if currentLine < 1 || currentColumn < 1 {
			break
		}
	}

	return knapsack, table[len(items)][knapSackCapacity], totalWeight
}
