package main

import (
	"fmt"
	"math"
	"testing"
)

func TestGenerateItems(t *testing.T) {
	for i := 0; i < 5; i++ {
		items := generateItems(int(math.Pow(10, float64(i+1))))
		writeItemsToFile("example_items/"+fmt.Sprint(i)+".data", items)
	}
}
