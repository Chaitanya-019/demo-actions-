package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"example/grpc_sql/sqlpb"
)

func main() {
	fmt.Println("Hello I'm Client!")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect : %v", err)
	}

	defer cc.Close()

	c := sqlpb.NewSqlServiceClient(cc)
	fmt.Println("Client created : ", c)
	getById(c)
}

func getById(c sqlpb.sqlServiceClient) {
	fmt.Println("Startingt the getByID function")
	req := &sqlpb.CompanyId{
		ID: "1",
	}

	res, _ := c.GetCompanyById(context.Background(), req)
	fmt.Println("Result from server : ", res)
}
