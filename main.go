package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	rand.Seed(time.Now().UnixNano())
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(1_000_000) + 1
	}
	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0 // или можно вернуть math.MinInt, но 0 для положительных чисел
	}
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) <= CHUNKS {
		return maximum(data)
	}
	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}
		wg.Add(1)
		go func(idx, s, e int) {
			defer wg.Done()
			maxValues[idx] = maximum(data[s:e])
		}(i, start, end)
	}
	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	startSingle := time.Now()
	maxSingle := maximum(data)
	elapsedSingle := time.Since(startSingle).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxSingle, elapsedSingle)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startMulti := time.Now()
	maxMulti := maxChunks(data)
	elapsedMulti := time.Since(startMulti).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxMulti, elapsedMulti)
}
