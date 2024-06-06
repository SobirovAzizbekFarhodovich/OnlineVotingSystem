package postgres

import (
	"database/sql"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	_ "github.com/lib/pq"

	g "vote/genproto"
)

func ConnectParty(t *testing.T) *PartyRepo {
	psql := "user=azizbek password=123 dbname=voting2 host=localhost port=5432 sslmode=disable"
	res, err := sql.Open("postgres", psql)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	return NewPartyRepo(res)
}

func CreateParty() *g.CreatePartyRequest {
	return &g.CreatePartyRequest{
		Id:          uuid.New().String(),
		Name:        "Diyorbek",
		Slogan:      "Isakov Uzb🇺🇿🇺🇿",
		OpenedDate:  "2005-01-01",
		Description: "Isakov Diyorbekki Qarindoshlariga",
	}
}

func TestCreateParty(t *testing.T) {
	db := ConnectParty(t)
	party := CreateParty()
	TestParty := g.CreatePartyRequest{
		Id:          party.Id,
		Name:        party.Name,
		Slogan:      party.Slogan,
		OpenedDate:  party.OpenedDate,
		Description: party.Description,
	}

	err := db.Create(&TestParty)
	if err != nil {
		t.Fatalf("Failed to create party: %v", err)
	}
	t.Log("party created successfully")

	idReq := g.ByIdPartyRequest{Id: party.Id}
	idRes, err := db.GetById(&idReq)
	if err != nil {
		t.Fatalf("Failed to get party: %v", err)
	}

	assert.Equal(t, party.Id, idRes.Id)
}

func TestDeleteParty(t *testing.T) {

	db := ConnectParty(t)
	party := CreateParty()
	TestParty := g.CreatePartyRequest{
		Id:          party.Id,
		Name:        party.Name,
		Slogan:      party.Slogan,
		OpenedDate:  party.OpenedDate,
		Description: party.Description,
	}

	err := db.Create(&TestParty)
	if err != nil {
		t.Fatalf("Error creating party: %v", err)
	}

	deleteReq := g.ByIdPartyRequest{Id: party.Id}
	_, err = db.Delete(&deleteReq)
	if err != nil {
		t.Fatalf("Error deleting party: %v", err)
	}

	_, err = db.GetById(&g.ByIdPartyRequest{Id: party.Id})
	if err == nil {
		t.Fatalf("Party should have been deleted")
	}
}

func TestGetAllParties(t *testing.T) {
	db := ConnectParty(t)
	party := CreateParty()
	TestParty := g.CreatePartyRequest{
		Id:          party.Id,
		Name:        party.Name,
		Slogan:      party.Slogan,
		OpenedDate:  party.OpenedDate,
		Description: party.Description,
	}

	err := db.Create(&TestParty)
	if err != nil {
		t.Fatalf("Error creating party: %v", err)
	}

	res, err := db.GetAll(&g.FilterPartyRequest{})
	if err != nil {
		t.Fatalf("Error getting all parties: %v", err)
	}

	assert.Equal(t, 1, len(res.Parties))
}
