package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Item struct {
	Index         int
	Weight        int
	Value         int
	ValueByWeight float64
}

func printItems(items []Item) {

	for i := 0; i < len(items); i++ {
		fmt.Printf("ITEM: %d VALOR: %d PESO: %d VALOR/PESO: %f\n", items[i].Index, items[i].Value, items[i].Weight, items[i].ValueByWeight)
	}

}

func readItemsFromFile(filePath string) []Item {
	var items []Item
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Failed to read file")
		panic(err)
	}

	fileLines := strings.Split(string(file), "\n")

	for i := 0; i < len(fileLines)-1; i++ {
		lineDatas := strings.Split(fileLines[i], " ")
		index, err := strconv.Atoi(lineDatas[0])
		if err != nil {
			fmt.Println("Failed to read file value")
			panic(err)
		}
		value, err := strconv.Atoi(lineDatas[1])
		if err != nil {
			fmt.Println("Failed to read file value")
			panic(err)
		}
		weight, err := strconv.Atoi(lineDatas[2])
		if err != nil {
			fmt.Println("Failed to read file value")
			panic(err)
		}
		items = append(items, Item{
			Index:         index,
			Value:         value,
			Weight:        weight,
			ValueByWeight: 0,
		})
	}
	return items
}

func writeItemsToFile(filePath string, items []Item) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Failed to create file")
		panic(err)
	}

	defer file.Close()

	for i := 0; i < len(items); i++ {
		line := fmt.Sprintf("%d %d %d 0\n", items[i].Index, items[i].Value, items[i].Weight)
		file.Write([]byte(line))
	}

	file.Sync()
}

func generateItems(numberOfItems, maxWeight, maxValue int) []Item {
	var items []Item
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numberOfItems; i++ {
		items = append(items, Item{
			Index:         i,
			Weight:        rand.Intn(maxWeight) + 1,
			Value:         rand.Intn(maxValue),
			ValueByWeight: 0,
		})
	}
	return items
}
