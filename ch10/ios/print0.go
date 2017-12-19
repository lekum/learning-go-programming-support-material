package main

import (
	"fmt"
	"os"
)

type metalloid struct {
	name   string
	number int32
	weight float64
}

func main() {

	var metalloids = []metalloid{
		metalloid{"Boron", 5, 10.81},
		metalloid{"Polonium", 84, 209.0},
	}

	file, err := os.Create("./metalloids.txt")
	if err != nil {
		fmt.Printf("Error creating metalloids file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	for _, m := range metalloids {
		fmt.Fprintf(
			file,
			"%-10s %-10d %-10.3f\n",
			m.name, m.number, m.weight,
		)
	}
}
