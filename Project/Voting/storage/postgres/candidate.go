package postgres

import (
	"database/sql"
	"fmt"
	g "root/genproto"

	"github.com/google/uuid"
)

type CandidateRepo struct {
	db *sql.DB
}

func NewCandidateRepo(db *sql.DB) *CandidateRepo {
	return &CandidateRepo{db: db}
}

func (c *CandidateRepo) Create(req *g.CreateCandidateRequest) (*g.Candidate_Void, error) {
	req.Id = uuid.NewString()
	_, err := c.db.Exec(`insert into candidate(id, election_id, party_id, public_id) 
						values($1, $2, $3, $4)`, req.Id, req.ElectionId, uuid.New().String(), uuid.New().String())
	if err != nil {
		return nil, fmt.Errorf("failed to create candidate: %v", err)
	}
	return &g.Candidate_Void{}, nil
}

func (c *CandidateRepo) GetByID(req *g.GetByIdCandidateRequest) (*g.GetByIdCandidateResponse, error) {
	var candidate g.GetAllCandidate
	candidate.Election= &g.Election{}
	err := c.db.QueryRow(`
	select c.id, c.party_id, c.public_id, e.name, e.date, e.id 
	from candidate c
	join election e on c.election_id = e.id 
	where c.id = $1`, req.Id).
	Scan(&candidate.Id, &candidate.PartyId, &candidate.PublicId, &candidate.Election.Name, &candidate.Election.Date, &candidate.Election.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get by id candidate: %v", err)
	}
	return &g.GetByIdCandidateResponse{Candidate: &candidate},  nil
}

func (c *CandidateRepo) GetAll(req *g.Candidate_Void) (*g.GetAllCandidatesResponse, error) {
	var candidates []*g.GetAllCandidate
	rows, err := c.db.Query(`
		SELECT c.id, c.party_id, c.public_id, e.name, e.date, e.id
		FROM candidate c
		JOIN election e ON c.election_id = e.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var candidate g.GetAllCandidate
		candidate.Election = &g.Election{}
		err := rows.Scan(&candidate.Id, &candidate.PartyId, &candidate.PublicId, &candidate.Election.Name, &candidate.Election.Date, &candidate.Election.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to get all candidates: %v", err)
		}
		candidates = append(candidates, &candidate)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return &g.GetAllCandidatesResponse{Candidates: candidates}, nil
}



func (c *CandidateRepo) Delete(req *g.DeleteCandidateRequest) (*g.Candidate_Void, error) {
	_, err := c.db.Exec(`delete from candidate where id = $1`, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete candidate: %v", err)
	}
	return &g.Candidate_Void{}, nil
}

func (c *CandidateRepo) Update(req *g.UpdateCandidateRequest) (*g.Candidate_Void, error) {
	_, err := c.db.Exec(`update candidate set election_id = $1, party_id = $2, public_id = $3 where id = $4`,
		req.Candidate.ElectionId, req.Candidate.PartyId, req.Candidate.PublicId, req.Candidate.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to update candidate: %v", err)
	}
	return &g.Candidate_Void{}, nil

}
