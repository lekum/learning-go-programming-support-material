package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink from its water bowl",
		"The dark bird of prey lands on a small tree after hunting for fish",
	}

	histogram := make(map[string]int)
	words := make(chan string)

	// splits and count
	go func(pipe chan<- string) {
		defer close(pipe)
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				pipe <- word
			}
		}
	}(words)

	// process word stream and count words
	// loops until wordsCh is closed
	for w := range words {
		histogram[w]++
	}

	for k, v := range histogram {
		fmt.Printf("%s\t(%d)\n", k, v)
	}
}
