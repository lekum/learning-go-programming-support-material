package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 1024 * 1024

type numbers [size]int

func initialize(nums *numbers) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(1000)
	}
}

func max(nums *numbers) int {
	temp := nums[0]
	for _, val := range nums {
		if val > temp {
			temp = val
		}
	}
	return temp
}

func main() {
	numbers1 := new(numbers)
	initialize(numbers1)
	fmt.Println("max:", max(numbers1))
}
