package postgres

import (
	"database/sql"
	pb "vote/genproto/genproto/voting"
	"github.com/google/uuid"
)

type PublicRepo struct {
	db *sql.DB
}

func NewPublicRepo(db *sql.DB) *PublicRepo {
	return &PublicRepo{db: db}
}
func (u *PublicRepo) Create(public *pb.CreatePublicRequest) (*pb.VoidPublicResponse, error) {
	public.Id = uuid.NewString()
	query :=
		`
	INSERT INTO public(id,first_name,last_name,birthday,gender,nation,party_id) VALUES ($1,$2,$3,$4,$5,$6,$7)
	`

	_, err := u.db.Exec(query, public.Id, public.FirstName, public.LastName,public.Birthday,public.Gender,public.Nation,public.PartyId)

	return nil, err
}

func (u *PublicRepo) GetById(id *pb.ByIdPublicRequest) (*pb.GetPublicResponse, error) {
	query :=
		`
	select id,first_name,last_name,birthday,gender,nation,party_id from public where id = $1
	`
	
	public := pb.GetPublicResponse{}
	err := u.db.QueryRow(query, id.GetId()).Scan(&public.Id, &public.FirstName, &public.LastName, &public.Birthday, &public.Gender, &public.Nation, &public.PartyId)

	if err != nil {
		return nil, err
	}

	return &public, nil
}

func (u *PublicRepo) GetAll(flt *pb.FilterPublicRequest) (*pb.GetAllPublicResponse, error) {

	query :=
		`
	select id, first_name, last_name, birthday, gender, nation, party_id from public 
	limit $1 offset $2
	`

	rows, err := u.db.Query(query, flt.GetLimit(), flt.GetOffset())

	if err != nil {
		return nil, err
	}

	publics := pb.GetAllPublicResponse{}
	for rows.Next() {
		public := pb.GetPublicResponse{}
		err := rows.Scan(&public.Id, &public.FirstName, &public.LastName,&public.Birthday,&public.Gender,&public.Nation,&public.PartyId)

		if err != nil {
			return nil, err
		}
		publics.Publics = append(publics.Publics, &public)
	}

	return &publics, nil
}

func (u *PublicRepo) Update(public *pb.GetPublicResponse)(*pb.VoidPublicResponse, error){
	query :=
		`
	update public set 
		first_name = $1,
		last_name = $2,
		birthday = $3,
		gender = $4,
		nation = $5,
		party_id = $6
	where 
		id = $7
	`

	_, err := u.db.Exec(query,public.GetFirstName(),public.GetLastName(),public.GetBirthday(),public.GetGender(),public.GetNation(),public.GetPartyId())

	return nil, err
}

func (u *PublicRepo) Delete(id *pb.ByIdPublicRequest)(*pb.VoidPublicResponse, error){
	query :=
		`
		delete from public where id = $1
		`

	_, err := u.db.Exec(query, id.GetId())

	return nil, err
}