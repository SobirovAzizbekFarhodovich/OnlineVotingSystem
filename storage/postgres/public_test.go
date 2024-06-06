package postgres

import (
	"database/sql"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	_ "github.com/lib/pq"

	g "vote/genproto/genproto/voting"
)

func ConnectPublic(t *testing.T) *PublicRepo {
	psql := "user=azizbek password=123 dbname=voting2 host=localhost port=5432 sslmode=disable"
	res, err := sql.Open("postgres", psql)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	return NewPublicRepo(res)
}

func CreatePublic() *g.CreatePublicRequest {
	return &g.CreatePublicRequest{
		Id:          uuid.New().String(),
		FirstName:   "Azizbek",
		LastName: 	 "Sobirov",
		Birthday: 	 "2007-07-25",
		Gender: 	 "m",
		Nation: 	 "Uzbekistan",
		PartyId: 	 " 1234567890",
	}
}

func TestCreatePublic(t *testing.T) {
	db := ConnectPublic(t)
	public := CreatePublic()
	TestPublic := g.CreatePublicRequest{
		Id:          public.Id,
		FirstName:   public.FirstName,
		LastName: 	 public.LastName,
		Birthday: 	 public.Birthday,
		Gender: 	 public.Gender,
		Nation: 	 public.Nation,
		PartyId: 	 public.PartyId,
	}

	_, err := db.Create(&TestPublic)
	if err != nil {
		t.Fatalf("Failed to create public: %v", err)
	}
	t.Log("public created successfully")

	idReq := g.ByIdPublicRequest{Id: public.Id}
	idRes, err := db.GetById(&idReq)
	if err != nil {
		t.Fatalf("Failed to get party: %v", err)
	}

	assert.Equal(t, public.Id, idRes.Id)
}

func TestDeletePublic(t *testing.T) {    

  db := ConnectPublic(t)
  public := CreatePublic()
  TestPublic := g.CreatePublicRequest{
    Id:         public.Id,
    FirstName:  public.FirstName,
    LastName:   public.LastName,
    Birthday:   public.Birthday,
    Gender:     public.Gender,
    Nation:     public.Nation,
    PartyId:    public.PartyId,
  }

  _, err := db.Create(&TestPublic)
  if err != nil {
    t.Fatalf("Error creating public: %v", err)
  }

  deleteReq := g.ByIdPublicRequest{Id: public.Id}
  _, err = db.Delete(&deleteReq)
  if err != nil {
    t.Fatalf("Error deleting public: %v", err)
  }

  _, err = db.GetById(&g.ByIdPublicRequest{Id: public.Id})
  if err == nil {
    t.Fatalf("Public should have been deleted")
  }
}

func TestGetAllPublics(t *testing.T) {
  db := ConnectPublic(t)
  public := CreatePublic()
  TestPublic := g.CreatePublicRequest{
    Id:         public.Id,
    FirstName:  public.FirstName,
	LastName:   public.LastName,
	Birthday:   public.Birthday,
	Gender:     public.Gender,
	Nation:     public.Nation,
	PartyId:    public.PartyId,
  }

  _, err := db.Create(&TestPublic)
  if err != nil {
    t.Fatalf("Error creating public: %v", err)
  }

  res, err := db.GetAll(&g.FilterPublicRequest{})
  if err != nil {
    t.Fatalf("Error getting all publics: %v", err)
  }

  assert.Equal(t, 1, len(res.Publics))
}