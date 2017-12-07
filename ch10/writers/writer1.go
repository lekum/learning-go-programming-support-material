package main

import (
	"fmt"
)

type channelWriter struct {
	Channel chan byte
}

func NewChannelWriter() *channelWriter {
	return &channelWriter{Channel: make(chan byte, 1024)}
}

func (cw *channelWriter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	go func() {
		defer close(cw.Channel)
		for _, b := range p {
			cw.Channel <- b
		}
	}()
	return len(p), nil
}

func main() {
	cw := NewChannelWriter()
	go func() {
		fmt.Fprintf(cw, "Stream me!")
	}()

	for c := range cw.Channel {
		fmt.Printf("%c\n", c)
	}
}
