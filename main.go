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
	// ваш код здесь
	if size <= 0 {
		return []int{}
	}
	randSource := rand.NewSource(time.Now().UnixNano())
	randRange := rand.New(randSource)
	slc := make([]int, size)
	for i := range size {
		slc[i] = randRange.Int()
	}
	return slc
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	maxNum := data[0]

	for _, v := range data {
		if maxNum < v {
			maxNum = v
		}
	}

	return maxNum
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	maxNums := make([]int, CHUNKS)
	size := len(data) / CHUNKS
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {

		start := i * size
		end := start + size

		if i == CHUNKS-1 {
			end = len(data)
		}
		
		wg.Add(1)
		go func (i int, slc []int)  {
			defer wg.Done()
			maxNums[i] = maximum(slc)
		}(i, data[start:end])
	}

	wg.Wait()

	return maximum(maxNums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	slc := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(slc)
	elapsed := 	int64(time.Since(start).Milliseconds())
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(slc)
	elapsed = int64((time.Since(start).Milliseconds()))

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
