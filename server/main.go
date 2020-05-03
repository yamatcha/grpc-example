package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/yamatcha/grpc-example/proto"
	"google.golang.org/grpc"
)

type customerService struct {
	customers []*pb.Person
	m         sync.Mutex
}

func (cs *customerService) ListPerson(p *pb.RequestType, stream pb.CustomerService_ListPersonServer) error {
	cs.m.Lock()
	defer cs.m.Unlock()
	for _, p := range cs.customers {
		if err := stream.Send(p); err != nil {
			return err
		}
	}
	return nil
}

func (cs *customerService) AddPerson(c context.Context, p *pb.Person) (*pb.ResponseType, error) {
	cs.m.Lock()
	defer cs.m.Unlock()
	cs.customers = append(cs.customers, p)
	return new(pb.ResponseType), nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:11111")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterCustomerServiceServer(server, new(customerService))
	log.Println("Server Started")
	server.Serve(lis)
}
