package main

import (
	"fmt"
	"log"
	"github.com/GO/49-dars/API_Gateway/api"
	"github.com/GO/49-dars/API_Gateway/api/handler"
	g "github.com/GO/49-dars/API_Gateway/genproto/voting"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main(){
	con, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	public := g.NewPublicServiceClient(con)
	party := g.NewPartyServiceClient(con)
	// /publiVote := genproto.NegwPublicVoteServiceClient(con)

	h := handler.NewHandler2(&public,&party)

	
	r := api.NewGin(h)
	fmt.Println("Server is running on port 8080")
	r.Run(":8080")

}