package main

import (
	"fmt"
	"sync"
	"time"
)

type Service struct {
	started bool
	stpCh   chan struct{}
	mutex   sync.RWMutex
	cache   map[int]string
}

func (s *Service) Start() {
	s.stpCh = make(chan struct{})
	go func() {
		s.mutex.Lock()
		s.started = true
		s.cache[1] = "Hello world"
		s.mutex.Unlock()
		<-s.stpCh // wait to be closed
	}()
}

func (s *Service) Stop() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.started {
		s.started = false
		close(s.stpCh)
	}
}

func (s *Service) Serve(id int) {
	s.mutex.RLock()
	msg := s.cache[id]
	s.mutex.RUnlock()
	if msg != "" {
		fmt.Println(msg)
	} else {
		fmt.Println("Hello, goodbye!")
	}
}

func main() {
	s := &Service{cache: make(map[int]string)}
	s.Start()
	time.Sleep(time.Second) // do some work
	s.Serve(1)
	s.Serve(2)
	s.Stop()
}
