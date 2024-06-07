package postgres


import (
	"database/sql"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	_ "github.com/lib/pq"

	g "root/genproto"
)


func ConnectCandidate(t *testing.T) *CandidateRepo {
	psql := "user=postgres password=20005 dbname=voting1 sslmode=disable"
	res, err := sql.Open("postgres", psql)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	return NewCandidateRepo(res)
}


func createCandidate() *g.Candidate {
	return &g.Candidate{
		Id:         uuid.New().String(),
		ElectionId:  "6ed0b65f-224a-4b09-a714-7b725d39994a",
		PartyId:    "6ed0b65f-224a-4b09-a714-7b725d3888b",
		PublicId:   uuid.New().String(),
	}
}


func TestCreateCandidate(t *testing.T) {
	db := ConnectCandidate(t)
	candidate := createCandidate()
	testCandidate := g.CreateCandidateRequest{
		Id:         candidate.Id,
		ElectionId: candidate.Id,
		PartyId:    candidate.PartyId,
		PublicId:   candidate.PublicId,
	}

	_, err := db.Create(&testCandidate)
	if err != nil {
		t.Fatalf("Failed to create candidate: %v", err)
	}
	t.Log("Candidate created successfully")

	idReq := g.GetByIdCandidateRequest{Id: candidate.Id}
	idRes, err := db.GetByID(&idReq)
	if err != nil {
		t.Fatalf("Failed to get candidate: %v", err)
	}

	assert.Equal(t, candidate.Id, idRes.Candidate.Id)
}

func TestDeleteCandidate(t *testing.T) {		

	db := ConnectCandidate(t)
	candidate := createCandidate()
	testCandidate := g.CreateCandidateRequest{
		Id:         candidate.Id,
		ElectionId: candidate.Id,
		PartyId:    candidate.PartyId,
		PublicId:   candidate.PublicId,
	}

	_, err := db.Create(&testCandidate)
	if err != nil {
		t.Fatalf("Error creating candidate: %v", err)
	}

	deleteReq := g.DeleteCandidateRequest{Id: candidate.Id}
	_, err = db.Delete(&deleteReq)
	if err != nil {
		t.Fatalf("Error deleting candidate: %v", err)
	}

	_, err = db.GetByID(&g.GetByIdCandidateRequest{Id: candidate.Id})
	if err == nil {
		t.Fatalf("Candidate should have been deleted")
	}
}

func TestGetAllCandidates(t *testing.T) {
	db := ConnectCandidate(t)
	candidate := createCandidate()
	testCandidate := g.CreateCandidateRequest{
		Id:         candidate.Id,
		ElectionId: candidate.Id,
		PartyId:    candidate.PartyId,
		PublicId:   candidate.PublicId,
	}

	_, err := db.Create(&testCandidate)
	if err != nil {
		t.Fatalf("Error creating candidate: %v", err)
	}

	res, err := db.GetAll(&g.Candidate_Void{})
	if err != nil {
		t.Fatalf("Error getting all candidates: %v", err)
	}

	assert.Equal(t, 1, len(res.Candidates))
}


func TestUpdateCandidate(t *testing.T) {
	db := ConnectCandidate(t)
	candidate := createCandidate()
	testCandidate := g.CreateCandidateRequest{
		Id:         candidate.Id,
		ElectionId: candidate.Id,
		PartyId:    candidate.PartyId,
		PublicId:   candidate.PublicId,
	}

	_, err := db.Create(&testCandidate)
	if err != nil {
		t.Fatalf("Error creating candidate: %v", err)
	}

	candidate.ElectionId = "3bfeb462-2146-4ea5-9c71-edfd3b81d673"
	candidate.PartyId = "3bfeb462-2146-4ea5-9c71-edfd3b81d673"
	candidate.PublicId = "3bfeb462-2146-4ea5-9c71-edfd3b81d673"

	_, err = db.Update(&g.UpdateCandidateRequest{Candidate: candidate})
	if err != nil {
		t.Fatalf("Error updating candidate: %v", err)
	}

	res, err := db.GetByID(&g.GetByIdCandidateRequest{Id: candidate.Id})
	if err != nil {
		t.Fatalf("Error getting candidate by id: %v", err)
	}

	assert.Equal(t, "3bfeb462-2146-4ea5-9c71-edfd3b81d673", res.Candidate.PartyId)
	assert.Equal(t, "3bfeb462-2146-4ea5-9c71-edfd3b81d673", res.Candidate.PublicId)
}
