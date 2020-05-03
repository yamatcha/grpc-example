package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/yamatcha/grpc-example/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func add(name string, age int) error {
	conn, err := grpc.Dial("127.0.0.1:11111", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewCustomerServiceClient(conn)

	person := &pb.Person{
		Name: name,
		Age:  int32(age),
	}
	_, err = client.AddPerson(context.Background(), person)
	return err
}

func list() error {
	conn, err := grpc.Dial("127.0.0.1:11111", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewCustomerServiceClient(conn)
	stream, err := client.ListPerson(context.Background(), new(pb.RequestType))
	if err != nil {
		return err
	}
	for {
		person, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(person)
	}
	return nil
}

func main() {
	err := add("list", 3)
	if err != nil {
		log.Fatal(err)
	}
	err = list()
	if err != nil {
		log.Fatal(err)
	}
}
