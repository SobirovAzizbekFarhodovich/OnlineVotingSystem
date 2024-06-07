package main

import (
	"fmt"
	"log"

	"root/api"
	_ "root/docs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ConnVoting, err := grpc.NewClient(":1212", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer ConnVoting.Close()

	ConnPublic, err := grpc.NewClient(":1313", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer ConnPublic.Close()

	r := api.NewGin(ConnVoting, ConnPublic)
	fmt.Println("Server is running on port 8080")
	r.Run(":8080")
}
