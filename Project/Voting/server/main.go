package main

import (
	"log"
	"net"
	g "root/genproto"
	"root/service"
	"root/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":1212")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	g.RegisterElectionServiceServer(s, service.NewElectionService(db))
	g.RegisterCandidateServiceServer(s, service.NewCandidateService(db))
	g.RegisterPublicVoteServiceServer(s, service.NewPublicVoteSerive(db))

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}
