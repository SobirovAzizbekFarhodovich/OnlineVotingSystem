package handler

import (
	"context"
	"net/http"
	g "root/genproto"

	"github.com/gin-gonic/gin"
)


// @Summary Create Election
// @Description Endpoint for creating a new Election
// @Name create_election
// @Tags Election
// @Accept json
// @Produce json
// @Param election body g.Election true "Election create"
// @Success 200 {object} string "Successfully created Election"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create Election"
// @Router /election [post]
func (h *Handler) CreateElection(c *gin.Context) {
	election := g.Election{}
	err := c.BindJSON(&election)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.Election.CreateElection(context.Background(), &g.CreateElectionRequest{Election: &election})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Election created successfully"})
}

// @Summary Get Elections
// @Description Endpoint for getting all Elections
// @Name get_election
// @Tags Election
// @Accept json
// @Produce json
// @Success 200 {object} g.GetAllElectionsResponse "Successfully fetched Elections"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get Elections"
// @Router /elections [get]
func (h *Handler) GetAllElections(c *gin.Context) {
	res, err := h.Election.GetAllElections(context.Background(), &g.Election_Void{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get Election by ID
// @Description Endpoint for getting an Election by ID
// @Name get_election_by_id
// @Tags Election
// @Accept json
// @Produce json
// @Param id path string true "Election ID"
// @Success 200 {object} g.GetElectionResponse "Successfully fetched Election"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to get Election"
// @Router /election/{id} [get]
func (h *Handler) GetByIdElection(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Election.GetByIdElection(context.Background(), &g.GetElectionRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update Election
// @Description Endpoint for updating an Election
// @Name update_election
// @Tags Election
// @Accept json
// @Produce json
// @Param election body g.UpdateElectionRequest true "Update Election"
// @Success 200 {object} string "Successfully updated Election"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to update Election"
// @Router /election [put]
func (h *Handler) UpdateElection(c *gin.Context) {
	election := g.Election{}
	err := c.BindJSON(&election)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.Election.UpdateElection(context.Background(), &g.UpdateElectionRequest{Election: &election})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Election updated successfully"})
}

// @Summary Delete Election
// @Description Endpoint for deleting an Election
// @Name delete_election
// @Tags Election
// @Accept json
// @Produce json
// @Param id path string true "Election ID"
// @Success 200 {object} string "Successfully deleted Election"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to delete Election"
// @Router /election/{id} [delete]
func (h *Handler) DeleteElection(c *gin.Context) {
	id := c.Param("id")
	_, err := h.Election.DeleteElection(context.Background(), &g.DeleteElectionRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Election deleted successfully"})
}