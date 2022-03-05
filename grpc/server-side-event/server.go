package main

import (
	"fmt"
	pb "go-examples/grpc/server-side-event/sse"
	"log"
	"sync"
	"time"
)

type server struct {
	pb.StreamServiceServer
}

func (s *server) FetchResponse(req *pb.Request, srv pb.StreamService_FetchResponseServer) error {
	log.Printf("fetch response for id: %d\n", req.GetId())

	wg := sync.WaitGroup{}
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(count int) {
			defer wg.Done()

			time.Sleep(time.Duration(count) * time.Second)
			resp := &pb.Response{Result: fmt.Sprintf("Request: #%d for ID: %d", count, req.GetId())}
			if err := srv.Send(resp); err != nil {
				log.Println(err)
			}
			log.Printf("finishing request number :%d", count)
		}(i)
	}

	wg.Wait()
	return nil
}
