package main

import (
	"booking-service-grpc/handler"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	s,_ := handler.NewServer()
	s.Start(&wg)
	wg.Wait()
}