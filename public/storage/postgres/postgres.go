package postgres

import (
	"database/sql"
	"fmt"
	cf "vote/config"
	"vote/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db      *sql.DB
	Publics storage.PublicI
	Parties storage.PartyI
}

func NewPostgresStorage(config cf.Config) (*Storage, error) {
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	pu := &PublicRepo{db}
	pa := &PartyRepo{db}
	Storage := &Storage{Db: db, Publics: pu, Parties: pa}
	return Storage, err
}

func  (s *Storage)Public() storage.PublicI {
	if s.Publics == nil {
		s.Publics = NewPublicRepo(s.Db)
	}
	return s.Publics
}
	func (s *Storage) Party() storage.PartyI{
		if s.Parties == nil {
			s.Parties = NewPartyRepo(s.Db)
		}
		return s.Parties
	}