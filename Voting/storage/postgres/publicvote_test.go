package postgres

import (
	"database/sql"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	_ "github.com/lib/pq"

	g "root/genproto"
)

func ConnectDatabase(t *testing.T) *PublicVote {
	psql := "user=postgres password=20005 dbname=voting1 sslmode=disable"
	db, err := sql.Open("postgres", psql)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	return NewPublicVote(db)
}

func createPublicVote() *g.PublicVote {
	return &g.PublicVote{
		Id:         uuid.New().String(),
		Election: &g.Election{Id: "6ed0b65f-224a-4b09-a714-7b725d39994a"},
		PublicId:   uuid.New().String(),
	}
}

func TestCreatePublicVote(t *testing.T) {
	db := ConnectDatabase(t)
	publicVote := createPublicVote()
	testPublicVote := g.CreatePublicVoteReq{
		PublicId:   publicVote.PublicId,
		ElectionId: publicVote.Election.Id,
		Id:         publicVote.Id,
	}

	_, err := db.Create(&testPublicVote)
	if err != nil {
		t.Fatalf("Failed to create public vote: %v", err)
	}
	t.Log("Public vote created successfully")

	idReq := g.GetPublicVoteRequest{Id: publicVote.Id}
	idRes, err := db.GetByPublicVoteId(&idReq)
	if err != nil {
		t.Fatalf("Failed to get public vote: %v", err)
	}

	assert.Equal(t, publicVote.Id, idRes.Id)
}

func TestDeletePublicVote(t *testing.T) {
	db := ConnectDatabase(t)
	publicVote := createPublicVote()
	testPublicVote := g.CreatePublicVoteReq{
		PublicId:   publicVote.PublicId,
		ElectionId: publicVote.Election.Id,
		Id:         publicVote.Id,
	}

	_, err := db.Create(&testPublicVote)
	if err != nil {
		t.Fatalf("Error creating public vote: %v", err)
	}

	_, err = db.Delete(&g.DeletePublicVoteRequest{Id: publicVote.Id})
	if err != nil {
		t.Fatalf("Error deleting public vote: %v", err)
	}

	_, err = db.GetByPublicVoteId(&g.GetPublicVoteRequest{Id: publicVote.Id})
	if err == nil {
		t.Fatalf("Public vote should be deleted")
	}
}

// func TestUpdatePublicVote(t *testing.T) {
// 	db := ConnectDatabase(t)
// 	publicVote := createPublicVote()
// 	testPublicVote := g.CreatePublicVoteReq{
// 		PublicId:   publicVote.PublicId,
// 		ElectionId: publicVote.ElectionId.Id,
// 		Id:         publicVote.Id,
// 	}

// 	_, err := db.Create(&testPublicVote)
// 	if err != nil {
// 		t.Fatalf("Error creating public vote: %v", err)
// 	}

// 	publicVote.ElectionId.Id = "3bfeb462-2146-4ea5-9c71-edfd3b81d673"
// 	publicVote.PublicId = "3bfeb462-2146-4ea5-9c71-edfd3b81d673"

// 	_, err = db.Update(&g.UpdatePublicVoteRequest{PublicVote: publicVote})
// 	if err != nil {
// 		t.Fatalf("Error updating public vote: %v", err)
// 	}

// 	res, err := db.GetByPublicVoteId(&g.GetPublicVoteRequest{Id: publicVote.Id})
// 	if err != nil {
// 		t.Fatalf("Error getting public vote by id: %v", err)
// 	}
// 	// createPublicVote creates a new PublicVote instance.
// 	assert.Equal(t, "3bfeb462-2146-4ea5-9c71-edfd3b81d673", res.ElectionId.Id)
// 	assert.Equal(t, "3bfeb462-2146-4ea5-9c71-edfd3b81d673", res.PublicId)
// }

func TestGetAllPublicVotes(t *testing.T) {
	db := ConnectDatabase(t)
	publicVote := createPublicVote()
	testPublicVote := g.CreatePublicVoteReq{
		PublicId:   publicVote.PublicId,
		ElectionId: publicVote.Election.Id,
		Id:         publicVote.Id,
	}

	_, err := db.Create(&testPublicVote)
	if err != nil {
		t.Fatalf("Error creating public vote: %v", err)
	}

	res, err := db.GetAll(&g.PublicVote_Void{})
	if err != nil {
		t.Fatalf("Error getting all public votes: %v", err)
	}

	assert.Equal(t, 1, len(res.PublicVote))
}
