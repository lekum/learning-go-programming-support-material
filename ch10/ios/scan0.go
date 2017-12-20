package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var name, hasRing string
	var diam, moons int

	data, err := os.Open("./planets.txt")
	if err != nil {
		fmt.Println("Unable to open planet data:", err)
		os.Exit(1)
	}
	defer data.Close()

	for {
		_, err := fmt.Fscanf(
			data,
			"%s %d %d %s\n",
			&name, &diam, &moons, &hasRing,
		)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Scan error:", err)
				os.Exit(1)
			}
		}
		fmt.Printf(
			"%-10s %-10d %-6d %-6s\n",
			name, diam, moons, hasRing,
		)

	}
}
