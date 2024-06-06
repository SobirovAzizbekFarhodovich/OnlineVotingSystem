package handler

import (
	g "root/genproto"

	"google.golang.org/grpc"
)



type Handler struct{
	Election g.ElectionServiceClient
	Candidate g.CandidateServiceClient
	PublicVote g.PublicVoteServiceClient
	Party g.PartyServiceClient
	Public g.PublicServiceClient
}

func NewHandler(voting, public *grpc.ClientConn) *Handler {
	return &Handler{
		Election: g.NewElectionServiceClient(voting),
		Candidate: g.NewCandidateServiceClient(voting),
		PublicVote: g.NewPublicVoteServiceClient(voting),
		Party: g.NewPartyServiceClient(public),
		Public: g.NewPublicServiceClient(public),
	}
}
