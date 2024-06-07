package handler

import (
	"context"
	"log"
	g "root/genproto"

	"github.com/gin-gonic/gin"
)


// @Summary Create Public Vote
// @Description Endpoint for creating a new Public Vote
// @Name create_public_vote
// @Tags PublicVote
// @Accept json
// @Produce json
// @Param publicvote body g.CreatePublicVoteReq true "Public Vote create"
// @Success 200 {object} string "Successfully created Public Vote"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create Public Vote"
// @Router /publicvote [post]
func(h *Handler) CreatePublicVote(c *gin.Context){
	publicvote := g.CreatePublicVoteReq{}
	err := c.BindJSON(&publicvote)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	_, err = h.PublicVote.CreatePublicVote(context.Background(), &publicvote)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

// @Summary Get Public Vote by ID
// @Description Endpoint for getting a Public Vote by ID
// @Name get_public_vote_by_id
// @Tags PublicVote
// @Accept json
// @Produce json
// @Param id path string true "Public Vote ID"
// @Success 200 {object} g.PublicVote "Successfully fetched Public Vote"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get Public Vote"
// @Router /publicvote/{id} [get]
func (h *Handler) GetByIdPublicVote(c *gin.Context){
	id := c.Param("id")
	log.Println(id)
	publicvote, err := h.PublicVote.GetByIdPublicVote(context.Background(), &g.GetPublicVoteRequest{Id: id})
	if err != nil{
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, publicvote)
}

// @Summary Get All Public Votes
// @Description Endpoint for getting all Public Votes
// @Name get_all_public_votes
// @Tags PublicVote
// @Accept json
// @Produce json
// @Success 200 {object} g.GetPublicVoteResponse "Successfully fetched Public Votes"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get Public Votes"
// @Router /publicvotes [get]
func (h *Handler) GetAllPublicVote(c *gin.Context){
	publicVotes, err := h.PublicVote.GetAllPublicVote(context.Background(), &g.PublicVote_Void{})
	if err != nil{
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}	
	c.JSON(200, publicVotes)
	}


// @Summary Update Public Vote
// @Description Endpoint for updating a Public Vote
// @Name update_public_vote
// @Tags PublicVote
// @Accept json
// @Produce json
// @Param publicvote body g.UpdatePublicVoteRequest true "Update Public Vote"
// @Success 200 {object} string "Successfully updated Public Vote"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to update Public Vote"
// @Router /publicvote [put]
func (h *Handler) UpdatePublicVote(c *gin.Context){
	publicvote := g.CreatePublicVoteReq{}
	err := c.BindJSON(&publicvote)
	if err != nil{
		c.JSON(400,  gin.H{"error": err.Error()})
		return
	}
	_, err = h.PublicVote.UpdatePublicVote(context.Background(), &g.UpdatePublicVoteRequest{PublicVote: &publicvote})
	if err != nil{
		c.JSON(500,  gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}


// @Summary Delete Public Vote
// @Description Endpoint for deleting a Public Vote
// @Name delete_public_vote
// @Tags PublicVote
// @Accept json
// @Produce json
// @Param id path string true "Public Vote ID"
// @Success 200 {object} string "Successfully deleted Public Vote"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to delete Public Vote"
// @Router /publicvote/{id} [delete]
func (h *Handler) DeletePublicVote(c *gin.Context){
	id := c.Param("id")
	_, err := h.PublicVote.DeletePublicVote(context.Background(), &g.DeletePublicVoteRequest{Id: id})
	if err != nil{
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}