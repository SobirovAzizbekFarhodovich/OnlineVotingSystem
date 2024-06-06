package handler

import g "github.com/GO/49-dars/API_Gateway/genproto"
import gv "github.com/GO/49-dars/API_Gateway/genproto/voting"



type Handler struct{
	Election g.ElectionServiceClient
	// Candidate g.CandidateServiceClient
	// Vote g.VoteServiceClient
	// PublicVote g.PublicVoteServiceClient
	Public gv.PublicServiceClient
	Party gv.PartyServiceClient
}

// func NewHandler(Election *g.ElectionServiceClient, PublicVote *g.PublicVoteServiceClient) *Handler {
// 	return &Handler{Election: *Election, PublicVote: *PublicVote}
// }

func NewHandler2(Public *gv.PublicServiceClient,Party *gv.PartyServiceClient) *Handler {
	return &Handler{Public: *Public,Party: *Party}
}