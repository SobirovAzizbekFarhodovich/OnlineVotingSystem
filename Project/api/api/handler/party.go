package handler

import (
	"context"
	"net/http"
	g "root/genproto"

	"github.com/gin-gonic/gin"
)

// @Summary Create Parties
// @Description Endpoint for creating a new party
// @Name create_party
// @Tags Party
// @Accept json
// @Produce json
// @Param party body g.CreatePartyRequest true "Party creation request payload"
// @Success 200 {object} string "Successfully created party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create party"
// @Router /party [post]
func (h *Handler) CreateParty(c *gin.Context) {
	party := g.CreatePartyRequest{}
	err := c.BindJSON(&party)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.Party.CreateParty(context.Background(), &party)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Party created successfully"})
}

// @Summary Get All Parties
// @Description Endpoint for fetching all parties
// @Name get_all_parties
// @Tags Party
// @Accept json
// @Produce json
// @Param filter query g.FilterPartyRequest true "Party GetAll request payload"
// @Success 200 {object} g.GetAllPartyResponse "Successfully fetched all parties"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to fetch all parties"
// @Router /parties [get]
func (h *Handler) GetAllParty(c *gin.Context) {
	res, err := h.Party.GetAllParties(context.Background(), &g.FilterPartyRequest{Limit: 10, Offset: 0})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get Party by ID
// @Description Endpoint for fetching a party by ID
// @Name get_party_by_id
// @Tags Party
// @Accept json
// @Produce json
// @Param id path string true "Party ID"
// @Success 200 {object} g.GetPartyResponse "Successfully fetched party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to fetch party"
// @Router /party/{id} [get]
func (h *Handler) GetByIdParty(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Party.GetPartyById(context.Background(), &g.ByIdPartyRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update Party
// @Description Endpoint for updating a party
// @Name update_party
// @Tags Party
// @Accept json
// @Produce json
// @Param party body g.UpdatePartyRequest true "Party Update request payload"
// @Success 200 {object} string "Successfully updated party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to update party"
// @Router /party [put]
func (h *Handler) UpdateParty(c *gin.Context){
	p := g.GetPartyResponse{}
	err := c.BindJSON(&p)
	_, err = h.Party.UpdateParty(context.Background(), &p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Party updated successfully"})
}


// @Summary Delete Party
// @Description Endpoint for deleting a party
// @Name delete_party
// @Tags Party
// @Accept json
// @Produce json
// @Param id path string true "Party ID"
// @Success 200 {object} string "Successfully deleted party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to delete party"
// @Router /party/{id} [delete]
func (h *Handler) DeleteParty(c *gin.Context) {
	id := c.Param("id")
	_, err := h.Party.DeleteParty(context.Background(), &g.ByIdPartyRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Party deleted successfully"})
}