package postgres

import (
	"database/sql"
	"root/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db         *sql.DB
	
	ElectionI  storage.Election
	CandidateI  storage.Candidate
	PublicVoteI       storage.PublicVote
	 
}

func ConnectDB() (*Storage, error) {
	psql := "user=postgres password=20005 dbname=voting1 sslmode=disable"
	db, err := sql.Open("postgres", psql)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	return &Storage{Db: db}, nil
}

func (s *Storage) Election() storage.Election  {
	if s.ElectionI == nil {
		s.ElectionI = &ElectionRepo{db: s.Db}
	}
	return s.ElectionI
}

func (s *Storage) Candidate() storage.Candidate  {
	if s.CandidateI == nil {
		s.CandidateI = &CandidateRepo{db: s.Db}
	}
	return s.CandidateI
}

func (s *Storage) PublicVote() storage.PublicVote  {
	if s.PublicVoteI == nil {
		s.PublicVoteI = &PublicVote{db: s.Db}
	}
	return s.PublicVoteI
}


