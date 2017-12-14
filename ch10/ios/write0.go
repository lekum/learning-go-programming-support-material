package main

import (
	"fmt"
	"os"
)

func main() {
	rows := []string{
		"The quick brown fox\n",
		"jumps over the lazy dog",
	}
	fout, err := os.Create("./filewrite.data")
	if err != nil {
		fmt.Printf("Error creating the file: %v\n", err)
		os.Exit(1)
	}
	defer fout.Close()

	for _, row := range rows {
		_, err = fout.WriteString(row)
		if err != nil {

			fmt.Printf("Error reading line: %v\n", err)
		}
	}
}
