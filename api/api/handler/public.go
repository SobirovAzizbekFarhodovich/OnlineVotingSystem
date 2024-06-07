package handler

import (
	"context"
	"net/http"
	g "root/genproto"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Public
// @Description Endpoint for creating a new Public
// @Name create_public
// @Balance create_public
// @Tags Public
// @Accept json
// @Produce json
// @Param public body g.CreatePublicRequest true "Public create"
// @Success 200 {object} string "Successfully created public"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to create public"
// @Router /public [POST]
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

// @Summary Get All Public
// @Description Get All Public
// @Name GetAllPublic
// @Balance GetAllPublic
// @Tags Public
// @Accept json
// @Produce json
// @Param public query g.FilterPublicRequest true "GetAllPublic"
// @Success 200 {object} g.GetAllPublicResponse "Successfully Get All Public"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to Get All public"
// @Router /publics [Get]
func (h *Handler) GetAllPublic(c *gin.Context) {
	filter := g.FilterPublicRequest{}
	limit := c.Query("limit")
	offset := c.Query("offset")
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)
	filter.Limit = int32(limitInt)
	filter.Offset = int32(offsetInt)
	res, err := h.Public.GetAllPublic(context.Background(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get Public
// @Description Get Public
// @Name GetPublic
// @Balance GetPublic
// @Tags Public
// @Accept json
// @Produce json
// @Param public body g.ByIdPublicRequest true "Get Public"
// @Success 200 {object} g.GetPublicResponse "Successfully Get Public"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to Get public"
// @Router /public/:id [Get]
func (h *Handler) GetByIdPublic(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Public.GetByIdPublic(context.Background(), &g.ByIdPublicRequest{Id: id})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}

// @Summary Update Public
// @Description Update Public
// @Name UpdatePublic
// @Balance UpdatePublic
// @Tags Public
// @Accept json
// @Produce json
// @Param public body g.GetPublicResponse true "UpdatePublic"
// @Success 200 {object} string "Successfully Update Public"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to Update public"
// @Router /public [Put]
func (h *Handler) UpdatePublic(c *gin.Context) {
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

// @Summary Delete Public
// @Description Delete Public
// @Name DeletePublic
// @Balance DeletePublic
// @Tags Public
// @Accept json
// @Produce json
// @Param public body g.ByIdPublicRequest true "DeletePublic"
// @Success 200 {object} string "Successfully Delete Public"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Failed to Delete public"
// @Router /public/:id [Delete]
func (h *Handler) DeletePublic(c *gin.Context) {
	id := c.Param("id")
	_, err := h.Public.DeletePublic(context.Background(), &g.ByIdPublicRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Public deleted successfully"})
}
