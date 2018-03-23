package main

import (
	"fmt"
	"math"
)

func main() {
	min := math.Inf(1)

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	for _, val := range x {
		tmp := float64(val)
		if tmp < min {
			min = tmp
		}
	}

	fmt.Println("Minimal array value:", min)
}
