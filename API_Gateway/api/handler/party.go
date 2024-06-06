package handler

import (
	"context"
	"net/http"
	g "github.com/GO/49-dars/API_Gateway/genproto/voting"

	"github.com/gin-gonic/gin"
)

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

func (h *Handler) GetAllParty(c *gin.Context) {
	res, err := h.Party.GetAllParties(context.Background(), &g.FilterPartyRequest{Limit: 10, Offset: 0})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetByIdParty(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Party.GetPartyById(context.Background(), &g.ByIdPartyRequest{Id: id})
	if err != nil{
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
	}


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

func (h *Handler) DeleteParty(c *gin.Context){
	id := c.Param("id")
	_, err := h.Party.DeleteParty(context.Background(), &g.ByIdPartyRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Party deleted successfully"})
}
