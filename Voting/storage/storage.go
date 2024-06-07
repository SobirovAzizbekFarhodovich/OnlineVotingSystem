package storage

import g "root/genproto"



type Candidate interface{
	Create(req *g.CreateCandidateRequest) (*g.Candidate_Void, error)
	Delete(req *g.DeleteCandidateRequest) (*g.Candidate_Void, error)
	GetAll(req *g.Candidate_Void) (*g.GetAllCandidatesResponse, error)
	GetByID(req *g.GetByIdCandidateRequest) (*g.GetByIdCandidateResponse, error)
	Update(req *g.UpdateCandidateRequest) (*g.Candidate_Void, error)
}

type Election interface{
 	Create(req *g.CreateElectionRequest) (*g.Election_Void, error)
 	Delete(req *g.DeleteElectionRequest) (*g.Election_Void, error)
 	GetAll(req *g.Election_Void) (*g.GetAllElectionsResponse, error)
 	GetByID(req *g.GetElectionRequest) (*g.GetElectionResponse, error)
 	Update(req *g.UpdateElectionRequest) (*g.Election_Void, error)
}


type PublicVote interface{
	Create(req *g.CreatePublicVoteReq) (*g.PublicVote_Void, error)
	Delete(req *g.DeletePublicVoteRequest) (*g.PublicVote_Void, error)
	GetAll(req *g.PublicVote_Void) (*g.GetPublicVoteResponse, error)
	GetByPublicVoteId(req *g.GetPublicVoteRequest) (*g.PublicVote, error)
	Update(req *g.UpdatePublicVoteRequest) (*g.PublicVote_Void, error)


}
