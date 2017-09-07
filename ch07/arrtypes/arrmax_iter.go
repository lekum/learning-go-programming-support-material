package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 1000

var nums [size]int

func init() {
	for i, _ := range nums {
		rand.Seed(time.Now().UnixNano())
		nums[i] = rand.Intn(1000)
	}
}

func max(nums [size]int) int {
	temp := nums[0]
	for _, val := range nums {
		if val > temp {
			temp = val
		}
	}
	return temp
}

func main() {
	fmt.Println("max:", max(nums))
}
