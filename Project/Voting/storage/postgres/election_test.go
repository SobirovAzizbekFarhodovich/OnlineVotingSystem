package postgres

import (
	"database/sql"
	g "root/genproto"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func databaseConnect(t *testing.T) *ElectionRepo {
	psql := "user=postgres password=20005 dbname=voting1 sslmode=disable"
	db, err := sql.Open("postgres", psql)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	return NewElectionRepo(db)
}

func createElection() *g.Election {
	return &g.Election{
		Id:   uuid.New().String(),
		Name: "OOO",
		Date: "2005-01-01",
	}
}

func TestCreateElection(t *testing.T) {
	db := databaseConnect(t)
	election := createElection()
	testElection := g.CreateElectionRequest{Election: election}
	_, err := db.Create(&testElection)
	if err != nil {
		t.Fatalf("Failed to create election: %v", err)
	}
	t.Log("Election created successfully")

	idReq := g.GetElectionRequest{Id: election.Id}
	idRes, err := db.GetByID(&idReq)
	if err != nil {
		t.Fatalf("Failed to get election: %v", err)
	}

	assert.Equal(t, election.Id, idRes.Election.Id)
}

func TestDeleteElection(t *testing.T) {
	db := databaseConnect(t)

	election := createElection()
	_, err := db.Create(&g.CreateElectionRequest{Election: election})
	if err != nil {
		t.Fatalf("Error creating election: %v", err)
	}

	_, err = db.Delete(&g.DeleteElectionRequest{Id: election.Id})
	if err != nil {
		t.Fatalf("Error deleting election: %v", err)
	}

	_, err = db.GetByID(&g.GetElectionRequest{Id: election.Id})
	if err == nil {
		t.Fatalf("Expected error 'election not found', but no error occurred")
	}	

	assert.Equal(t, "election not found: sql: no rows in result set", err.Error())
}


func TestUpdateElection(t *testing.T) {
	db := databaseConnect(t)

	election := createElection()
	id := "3bfeb462-2146-4ea5-9c71-edfd3b81d673"
	election.Id = id

	_, err := db.Update(&g.UpdateElectionRequest{Election: election})
	if err != nil {
		t.Fatalf("Error deleting election: %v", err)
	}

	res, err := db.GetByID(&g.GetElectionRequest{Id: election.Id})
	if err != nil {
		t.Fatalf("Expected error 'election not found', but no error occurred")
	}

	assert.Equal(t, "succuses", res)
}


func TestGetAllElections(t *testing.T) {
	db := databaseConnect(t)

	election := createElection()
	_, err := db.Create(&g.CreateElectionRequest{Election: election})
	if err != nil {
		t.Fatalf("Error creating election: %v", err)
	}

	res, err := db.GetAll(&g.Election_Void{})
	if err != nil {
		t.Fatalf("Error getting all elections: %v", err)
	}

	assert.Equal(t, 1, len(res.Elections))
}

func TestGetElectionById(t *testing.T) {
	db := databaseConnect(t)

	election := createElection()
	_, err := db.Create(&g.CreateElectionRequest{Election: election})
	if err != nil {
		t.Fatalf("Error creating election: %v", err)
	}

	res, err := db.GetByID(&g.GetElectionRequest{Id: election.Id})
	if err != nil {
		t.Fatalf("Error getting election by id: %v", err)
	}

	assert.Equal(t, election.Id, res.Election.Id)
}