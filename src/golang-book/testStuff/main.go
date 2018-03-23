package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const SYMBOLS string = "abcdefghijklmnopqrstuvwxyz0123456789"

func generateString(a uint) (float64, string) {
	rand.Seed(time.Now().UnixNano())
	strArray := make([]string, a)

	startTime := time.Now().UnixNano()
	for i := uint(0); i < a; i++ {
		symbol := string(SYMBOLS[rand.Intn(36)])
		strArray[i] = symbol
	}
	str := strings.Join(strArray, "")
	finishTime := time.Now().UnixNano()

	return float64(finishTime-startTime) / 1000000000, str
}

func main() {
	times := 100
	elapsedSlice := make([]float64, times)
	for i := 0; i < times; i++ {
		elapsed, _ := generateString(10000000)
		elapsedSlice[i] = elapsed
	}

	elapsedAverage := func() float64 {
		var sum float64
		for _, value := range elapsedSlice {
			sum += value
		}
		return sum / float64(times)
	}()
	fmt.Println("Average execution time:", elapsedAverage)
}
