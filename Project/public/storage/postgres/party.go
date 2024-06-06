package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	pb "vote/genproto"

	"github.com/google/uuid"
)

type PartyRepo struct {
	db *sql.DB
}

func NewPartyRepo(db *sql.DB) *PartyRepo {
	return &PartyRepo{db: db}
}
func (u *PartyRepo) Create(party *pb.CreatePartyRequest) error {
	party.Id = uuid.NewString()
	query :=
		`
	INSERT INTO party(id,name,slogan,opened_date,description) VALUES ($1,$2,$3,$4,$5)
	`
	date, err := time.Parse("2006-01-02", party.OpenedDate)
	if err != nil {
		return fmt.Errorf("invalide %v", err)
	}

	_, err = u.db.Exec(query, party.Id, party.Name, party.Slogan, date, party.Description)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *PartyRepo) GetById(id *pb.ByIdPartyRequest) (*pb.GetPartyResponse, error) {
	query :=
		`
	select id,name,slogan,opened_date,description from party where id = $1
	`

	party := pb.GetPartyResponse{}
	err := u.db.QueryRow(query, id.GetId()).Scan(&party.Id, &party.Name, &party.Slogan, &party.OpenedDate, &party.Description)

	if err != nil {
		return nil, err
	}

	return &party, nil
}

func (u *PartyRepo) GetAll(flt *pb.FilterPartyRequest) (*pb.GetAllPartyResponse, error) {

	query :=
		`
	select id, name, slogan, opened_date,description from party 
	limit $1 offset $2
	`

	rows, err := u.db.Query(query, flt.GetLimit(), flt.GetOffset())

	if err != nil {
		return nil, err
	}

	parties := pb.GetAllPartyResponse{}
	for rows.Next() {
		party := pb.GetPartyResponse{}
		err := rows.Scan(&party.Id, &party.Name, &party.Slogan, &party.OpenedDate, &party.Description)

		if err != nil {
			return nil, err
		}
		parties.Parties = append(parties.Parties, &party)
	}

	return &parties, nil
}

func (u *PartyRepo) Update(party *pb.GetPartyResponse) (*pb.VoidPartyResponse, error) {
	query :=
		`
	update party set 
		name = $1,
		slogan = $2,
		opened_date = $3,
		description = $4
	where 
		id = $5
	`
	date, err := time.Parse("2006-01-02", party.OpenedDate)
	if err != nil {
		return nil, fmt.Errorf("invalide %v", err)
	}
	
	_, err = u.db.Exec(query, party.GetName(), party.GetSlogan(), date, party.GetDescription(), party.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.VoidPartyResponse{}, nil
}

func (u *PartyRepo) Delete(id *pb.ByIdPartyRequest) (*pb.VoidPartyResponse, error) {
	query :=
		`
		delete from party where id = $1
		`

	_, err := u.db.Exec(query, id.GetId())
	if err != nil {
		return nil,err
	}
	return &pb.VoidPartyResponse{}, nil
}
