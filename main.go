package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("hello world")
	start := time.Now()
	fmt.Println(calculatePi(1000000000, 4))
	elapsed := time.Since(start)
	fmt.Println("calculatePi async took %s", elapsed)

	start = time.Now()
	fmt.Println(calculatePiSync(1000000000))
	elapsed = time.Since(start)
	fmt.Println("calculatePi sync took %s", elapsed)
}

func calculatePiSync(numberOfIterations int) float64 {
	// Pi/4 = 1 - 1/3 + 1/5 - 1/7 + ...
	var pi float64 = 1
	pi += calculatePart(0, numberOfIterations)

	return pi * 4
}

func calculatePi(numberOfIterations int, numberOfBatches int) float64 {
	// Pi/4 = 1 - 1/3 + 1/5 - 1/7 + ...
	var wg sync.WaitGroup
	var pi float64 = 1
	var batches = calculateBatches(numberOfIterations, numberOfBatches)
	channel := make(chan float64, len(batches))

	for _, batch := range batches {
		wg.Add(1)
		go func(b Batch) {
			defer wg.Done()
			channel <- calculatePart(b.from, b.to)
			fmt.Println("calc done")
		}(batch)
	}

	fmt.Println("before wait")

	go func() {
		wg.Wait()
		close(channel)
		fmt.Println("close")
	}()

	fmt.Println("after wait")

	for s := range channel {
		fmt.Println("waiting for channel")
		pi += s
	}

	return pi * 4
}

type Batch struct {
	from int
	to   int
}

func calculateBatches(numberOfIterations int, numberOfBatches int) []Batch {
	var batchSize int = int(numberOfIterations / numberOfBatches)
	var batches []Batch = make([]Batch, 0)
	for i := 0; i < numberOfBatches; i++ {
		var newBatch = Batch{from: i * batchSize, to: (i+1)*batchSize - 1}
		batches = append(batches, newBatch)
	}
	return batches
}

func calculatePart(from int, to int) float64 {
	var pi float64 = 0
	for i := from + 1; i <= to; i++ {
		var isNegative bool = true
		if i%2 == 0 {
			isNegative = false
		}
		var factor float64 = float64(i*2 + 1)
		if isNegative {
			pi -= 1 / factor
		} else {
			pi += 1 / factor
		}
		// isNegative = !isNegative
	}
	return pi
}
