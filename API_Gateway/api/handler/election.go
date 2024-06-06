package handler

import (
	"context"
	"net/http"
	g "github.com/GO/49-dars/API_Gateway/genproto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateElections(c *gin.Context) {
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

func (h *Handler) GetAllElections(c *gin.Context) {
	res, err := h.Election.GetAllElections(context.Background(), &g.Election_Void{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetByIdElection(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Election.GetByIdElection(context.Background(), &g.GetElectionRequest{Id: id})
	if err != nil{
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
	}


func (h *Handler) UpdateElection(c *gin.Context){
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

func (h *Handler) DeleteElection(c *gin.Context){
	id := c.Param("id")
	_, err := h.Election.DeleteElection(context.Background(), &g.DeleteElectionRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Election deleted successfully"})
}


