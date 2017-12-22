package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rows := []string{
		"The quick brown fox\n",
		"jumps over the lazy dog\n",
	}
	fout, err := os.Create("./filewrite.data")
	if err != nil {
		fmt.Printf("Error creating the file: %v\n", err)
		os.Exit(1)
	}
	defer fout.Close()
	writer := bufio.NewWriter(fout)

	for _, row := range rows {
		_, err = writer.WriteString(row)
		if err != nil {
			fmt.Printf("Error writing line: %v\n", err)
		}
	}
	writer.Flush()
}
