package main

import (
	"log"
	"net"
	cf "vote/config"
	pb "vote/genproto/genproto/voting"
	"vote/service"
	"vote/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	config := cf.Load()

	db, err := postgres.NewPostgresStorage(config)
	if err != nil {
		log.Fatal(err)
	}

	liss, err := net.Listen("tcp", config.TCP_PORT)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterPublicServiceServer(s, service.NewPublicService(db))
	pb.RegisterPartyServiceServer(s, service.NewPartyService(db))

	log.Println("Server is listening", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatal(err)
	}
}
