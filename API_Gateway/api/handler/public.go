package handler

import (
	"context"
	"net/http"
	g "github.com/GO/49-dars/API_Gateway/genproto/voting"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePublic(c *gin.Context) {
	public := g.CreatePublicRequest{}
	err := c.BindJSON(&public)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	_, err = h.Public.CreatePublic(context.Background(), &public)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Public created successfully"})

}

func (h *Handler) GetAllPublic(c *gin.Context) {
	res, err := h.Public.GetAllPublic(context.Background(), &g.FilterPublicRequest{Limit: 10, Offset: 0})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetByIdPublic(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Public.GetByIdPublic(context.Background(), &g.ByIdPublicRequest{Id: id})
	if err != nil{
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
	}


func (h *Handler) UpdatePublic(c *gin.Context){
	public := g.GetPublicResponse{}
	err := c.BindJSON(&public)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = h.Public.UpdatePublic(context.Background(), &public)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Public updated successfully"})
}

func (h *Handler) DeletePublic(c *gin.Context){
	id := c.Param("id")
	_, err := h.Public.DeletePublic(context.Background(), &g.ByIdPublicRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Public deleted successfully"})
}


