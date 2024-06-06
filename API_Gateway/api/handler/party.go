package handler

import (
	"context"
	"net/http"
	g "github.com/GO/49-dars/API_Gateway/genproto/voting"

	"github.com/gin-gonic/gin"
)
// @Summary Create Parties
// @Description Endpoint for creating a new parties
// @Name create_party
// @Balance create_party
// @Tags Party
// @Accept json
// @Produce json
// @Param party body g.CreatePartyRequest true "Party creation request payload"
// @Success 200 {object} string "Successfully created party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create party"
// @Router /party [POST]
func (h *Handler) CreatePartys(c *gin.Context) {
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
// @Summary GetAllParties
// @Description GetAll
// @Name GetallParties
// @Balance GetallParties
// @Tags Party
// @Accept json
// @Produce json
// @Param party body g.FilterPartyRequest true "Party GetAll request payload"
// @Success 200 {object} g.GetAllPartyResponse "Successfully Get All party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to Get All party"
// @Router /party [GET]
func (h *Handler) GetAllParty(c *gin.Context) {
	res, err := h.Party.GetAllParties(context.Background(), &g.FilterPartyRequest{Limit: 10, Offset: 0})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
// @Summary GetByIdParty
// @Description GetAll
// @Name GetByIdParty
// @Balance GetByIdParty
// @Tags Party
// @Accept json
// @Produce json
// @Param party body g.ByIdPartyRequest true "Party GetById request payload"
// @Success 200 {object} g.GetPartyResponse "Successfully Get All party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to Get All party"
// @Router /party/:id [GET]
func (h *Handler) GetByIdParty(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Party.GetPartyById(context.Background(), &g.ByIdPartyRequest{Id: id})
	if err != nil{
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
	}

// @Summary Update Party
// @Description Update Party elements
// @Name Update Party
// @Balance Update Party
// @Tags Party
// @Accept json
// @Produce json
// @Param party body g.GetPartyResponse true "Party Update request payload"
// @Success 200 {object} string "Successfully Update party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to Update party"
// @Router /party [Put]
func (h *Handler) UpdateParty(c *gin.Context){
	party := g.GetPartyResponse{}
	err := c.BindJSON(&party)
	_, err = h.Party.UpdateParty(context.Background(), &party)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Party updated successfully"})
}
// @Summary Delete Party
// @Description Delete Party elements
// @Name Delete Party
// @Balance Delete Party
// @Tags Party
// @Accept json
// @Produce json
// @Param party body g.ByIdPartyRequest true "Party Delete request payload"
// @Success 200 {object} string "Successfully Delete party"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to Delete party"
// @Router /party/:id [Delete]
func (h *Handler) DeleteParty(c *gin.Context){
	id := c.Param("id")
	_, err := h.Party.DeleteParty(context.Background(), &g.ByIdPartyRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Party deleted successfully"})
}
