package main

import (
	"fmt"
	"runtime"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	totalMemKb := bToKb(m.TotalAlloc)
	totalMemMB := bToMb(m.TotalAlloc)
	fmt.Println(totalMemKb, " Kb\n", totalMemMB, " Mb")
}

func bToKb(b uint64) uint64 {
	return b / 1024
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
