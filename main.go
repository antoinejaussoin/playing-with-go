package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world")
	start := time.Now()
	fmt.Println(calculatePi(1000000000, 80))
	elapsed := time.Since(start)
	fmt.Println("calculatePi took %s", elapsed)
}

func calculatePi(numberOfIterations int, numberOfBatches int) float64 {
	// Pi/4 = 1 - 1/3 + 1/5 - 1/7 + ...
	var pi float64 = 1
	var batches = calculateBatches(numberOfIterations, numberOfBatches)

	for _, batch := range batches {
		pi += calculatePart(batch.from, batch.to)
	}
	//for p:= 0; p < 8; p++ {

	//}

	// for i := 1; i <= numberOfIterations; i++ {
	// 	var isNegative bool = true;
	// 	if i % 2 == 0 {
	// 		isNegative = false
	// 	}
	// 	var factor float64 = float64(i * 2 + 1)
	// 	if isNegative {
	// 		pi -= 1 / factor
	// 	} else {
	// 		pi += 1 / factor
	// 	}
	// }
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
	}
	return pi
}
