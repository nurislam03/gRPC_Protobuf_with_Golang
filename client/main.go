package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nurislam03/gRPC_Protobuf_with_Golang/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	a := 10
	b := 5
	req := &proto.Request{A: int64(a), B: int64(b)}

	ctx := context.Background()

	response, err := client.Add(ctx, req)
	if err != nil {
		erro := err.Error()
		log.Println(erro)
		return
	}

	result := fmt.Sprint(response.Result)
	log.Println(result)

	// subtract
	response, err = client.Subtract(ctx, req)
	if err != nil {
		erro := err.Error()
		log.Println(erro)
		return
	}

	result = fmt.Sprint(response.Result)
	log.Println(result)
}
