package handler

import (
	"context"
	g "root/genproto"

	"github.com/gin-gonic/gin"
)

// @Summary Create Candidate
// @Description Endpoint for creating a new Candidate
// @Name create_candidate
// @Tags Candidate
// @Accept json
// @Produce json
// @Param candidate body g.CreateCandidateRequest true "Candidate create"
// @Success 200 {object} string "Successfully created Candidate"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create Candidate"
// @Router /candidate [post]
func (h *Handler) CreateCandidate(c *gin.Context) {
	candidate := g.CreateCandidateRequest{}
	err := c.BindJSON(&candidate)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err = h.Candidate.CreateCandidate(context.Background(), &candidate)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Candidate created successfully"})
}

// @Summary Get Candidate by ID
// @Description Endpoint for getting a Candidate by ID
// @Name get_candidate_by_id
// @Tags Candidate
// @Accept json
// @Produce json
// @Param id path string true "Candidate ID"
// @Success 200 {object} g.GetByIdCandidateResponse "Successfully fetched Candidate"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get Candidate"
// @Router /candidate/{id} [get]
func (h *Handler) GetByIDCandidate(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Candidate.GetByIdCandidate(context.Background(), &g.GetByIdCandidateRequest{Id: id})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)

}

// @Summary Delete Candidate
// @Description Endpoint for deleting a Candidate
// @Name delete_candidate
// @Tags Candidate
// @Accept json
// @Produce json
// @Param id path string true "Candidate ID"
// @Success 200 {object} string "Successfully deleted Candidate"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to delete Candidate"
// @Router /candidate/{id} [delete]
func (h *Handler) DeleteCandidate(c *gin.Context) {
	id := c.Param("id")
	_, err := h.Candidate.DeleteCandidate(context.Background(), &g.DeleteCandidateRequest{Id: id})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Candidate deleted successfully"})

}

// @Summary Update Candidate
// @Description Endpoint for updating a Candidate
// @Name update_candidate
// @Tags Candidate
// @Accept json
// @Produce json
// @Param candidate body g.UpdateCandidateRequest true "Update Candidate"
// @Success 200 {object} string "Successfully updated Candidate"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to update Candidate"
// @Router /candidate [put]
func (h *Handler) UpdateCandidate(c *gin.Context) {
	candidate := g.UpdateCandidateRequest{}
	err := c.BindJSON(&candidate)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err = h.Candidate.UpdateCandidate(context.Background(), &g.UpdateCandidateRequest{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Candidate updated successfully"})
}

// @Summary Get All Candidates
// @Description Endpoint for getting all Candidates
// @Name get_all_candidates
// @Tags Candidate
// @Accept json
// @Produce json
// @Success 200 {object} g.GetAllCandidatesResponse "Successfully fetched Candidates"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get Candidates"
// @Router /candidates [get]
func (h *Handler) GetAllCandidaties(c *gin.Context) {
	res, err := h.Candidate.GetAllCandidates(context.Background(), &g.Candidate_Void{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}
