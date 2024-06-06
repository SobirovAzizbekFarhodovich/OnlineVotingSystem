package postgres

import (
	"database/sql"
	"fmt"
	g "root/genproto"

	"github.com/google/uuid"
)

type PublicVote struct {
	db *sql.DB
}

func NewPublicVote(db *sql.DB) *PublicVote {
	return &PublicVote{db: db}
}

func (pv *PublicVote) Create(req *g.CreatePublicVoteReq) (*g.PublicVote_Void, error) {
	tran, err := pv.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tran.Commit()

	req.Id = uuid.NewString()
	_, err = pv.db.Exec("insert into public_vote(id, election_id, public_id) values($1, $2, $3)",
		req.Id, req.ElectionId, req.PublicId)
	if err != nil {
		return nil, fmt.Errorf("failed to create public vote: %v", err)
	}

	id := uuid.NewString()
	_, err = pv.db.Exec("insert into votes(id, candidate_id) values($1, $2)", id, req.CandidateId)
	if err != nil {
		return nil, fmt.Errorf("failed to create votes: %v", err)
	}
	
	return &g.PublicVote_Void{}, nil
		
}

func (pv *PublicVote) GetByPublicVoteId(req *g.GetPublicVoteRequest) (*g.PublicVote, error) {
	var publicVote g.PublicVote
	publicVote.Election = &g.Election{} 
	err := pv.db.QueryRow(`
	SELECT p.id, p.public_id, e.id as election, e.name
	FROM public_vote p
	JOIN election e ON p.election_id = e.id 
	WHERE p.id = $1
	`, req.Id).
		Scan(&publicVote.Id, &publicVote.PublicId, &publicVote.Election.Id, &publicVote.Election.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get public vote by public id: %v", err)
	}

	return &g.PublicVote{
		Id:         publicVote.Id,
		PublicId:   publicVote.PublicId,
		Election: &g.Election{Id: publicVote.Election.Id, Name: publicVote.Election.Name},
	}, nil
}

func (pv *PublicVote) Delete(req *g.DeletePublicVoteRequest) (*g.PublicVote_Void, error) {
	_, err := pv.db.Exec("delete from public_vote where id = $1", req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete public vote: %v", err)
	}
	return &g.PublicVote_Void{}, nil
}

func (pv *PublicVote) Update(req *g.UpdatePublicVoteRequest) (*g.PublicVote_Void, error) {
	_, err := pv.db.Exec("update public_vote set election_id = $1, public_id = $2 where id = $3",
		req.PublicVote.Id, req.PublicVote.PublicId, req.PublicVote.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to update public vote: %v", err)
	}
	return &g.PublicVote_Void{}, nil
}

func (pv *PublicVote) GetAll(req *g.PublicVote_Void) (*g.GetPublicVoteResponse, error) {
	rows, err := pv.db.Query(
	`SELECT p.id, p.public_id, e.id as election, e.name
	FROM public_vote p
	JOIN election e ON p.election_id = e.id `)
	if err != nil {
		return nil, fmt.Errorf("failed to get all public votes: %v", err)
	}
	defer rows.Close()

	var publicVotes []*g.PublicVote
	for rows.Next() {
		var publicVote g.PublicVote
		publicVote.Election = &g.Election{}
		err := rows.Scan(&publicVote.Id, &publicVote.PublicId, &publicVote.Election.Id, &publicVote.Election.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan public vote: %v", err)
		}
		publicVotes = append(publicVotes, &publicVote)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return &g.GetPublicVoteResponse{PublicVote: publicVotes}, nil
}
