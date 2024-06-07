package postgres

import (
	"database/sql"
	"fmt"
	g "root/genproto"
	"time"

	"github.com/google/uuid"
)

type ElectionRepo struct {
	db *sql.DB
}

func NewElectionRepo(db *sql.DB) *ElectionRepo {
	return &ElectionRepo{db: db}
}

func (e *ElectionRepo) Create(req *g.CreateElectionRequest) (*g.Election_Void, error) {
	if req.Election.Id == "" {
		req.Election.Id = uuid.NewString()
	}

	date, err := time.Parse("2006-01-02", req.Election.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %v", err)
	}

	_, err = e.db.Exec("INSERT INTO election(id, name, date) VALUES($1, $2, $3)",
		req.Election.Id, req.Election.Name, date)
	if err != nil {
		return nil, fmt.Errorf("failed to create election: %v", err)
	}
	return &g.Election_Void{}, nil
}

func (e *ElectionRepo) GetAll(req *g.Election_Void) (*g.GetAllElectionsResponse, error) {
	rows, err := e.db.Query("select id, name, date from election")
	if err != nil {
		return nil, fmt.Errorf("failed to get all elections: %v", err)
	}
	defer rows.Close()
	var elections []*g.Election
	for rows.Next() {
		var election g.Election
		err := rows.Scan(&election.Id, &election.Name, &election.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to scan election: %v", err)
		}
		elections = append(elections, &election)

	}
	return &g.GetAllElectionsResponse{Elections: elections}, nil
}

func (e *ElectionRepo) GetByID(req *g.GetElectionRequest) (*g.GetElectionResponse, error) {
	var election g.Election
	err := e.db.QueryRow("select id, name, date from election where id = $1", req.Id).
		Scan(&election.Id, &election.Name, &election.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to get election by id: %v", err)
	}
	return &g.GetElectionResponse{Election: &election}, nil
}

func (e *ElectionRepo) Update(req *g.UpdateElectionRequest) (*g.Election_Void, error) {
	date, err := time.Parse("2006-01-02", req.Election.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %v", err)
	}
	_, err = e.db.Exec("update election set name = $1, date = $2 where id = $3",
		req.Election.Name, date, req.Election.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to update election: %v", err)
	}
	return &g.Election_Void{}, nil
}

func (e *ElectionRepo) Delete(req *g.DeleteElectionRequest) (*g.Election_Void, error) {
	_, err := e.db.Exec("delete from election where id = $1", req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete election: %v", err)
	}
	return &g.Election_Void{}, nil
}
