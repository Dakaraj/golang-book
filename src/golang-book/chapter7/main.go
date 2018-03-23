package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(args ...float64) float64 {
	summed := 0.0
	for _, val := range args {
		summed += val
	}
	return summed
}

func inpuHandler() []float64 {
	var collection []float64
	var tempString string
	reader := bufio.NewScanner(os.Stdin)
	for tempString != "exit" {
		fmt.Print("\nInput a float number different from 0:\n>")
		reader.Scan()
		tempString = reader.Text()

		if value, err := strconv.ParseFloat(tempString, 64); value != 0 && err == nil {
			collection = append(collection, value)
			fmt.Println(collection)
		}
	}

	return collection
}

func isOdd(i int64) (int64, bool) {
	rem := i % 2
	if rem == 1 {
		return rem, true
	}
	return rem, false
}

func max(arr []float64) float64 {
	var mx float64
	for _, val := range arr {
		if val > mx {
			mx = val
		}
	}

	return mx
}

func evenGenerator() func() int64 {
	var even int64
	return func() int64 {
		rt := even
		even += 2
		return rt
	}
}

func fib(i uint) uint {
	if i == 0 || i == 1 {
		return 1
	}
	return i + fib(i-1)
}

func main() {
	fmt.Println(fib(3))
	// nextEven := evenGenerator()
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())
	// // fmt.Println(isOdd(11))
	// fmt.Println(max([]float64{12.6, 14.5, 22.7, 13.1}))
	// tmp := inpuHandler()
	// fmt.Println(sum(tmp...))
}
