package handler

// import (
// 	"context"
// 	g "github.com/GO/49-dars/API_Gateway/genproto"

// 	"github.com/gin-gonic/gin"
// )



// func(h *Handler) CreatePublicVote(c *gin.Context){
// 	publicvote := g.CreatePublicVoteReq{}
// 	err := c.BindJSON(&publicvote)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}
	
// 	_, err = h.PublicVote.CreatePublicVote(context.Background(), &publicvote)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, gin.H{"message": "success"})
// }

// func (h *Handler) GetPublicVoteById(c *gin.Context){
// 	id := g.GetPublicVoteRequest{Id: c.Param(":id")}
// 	publicvote, err := h.PublicVote.GetByIdPublicVote(context.Background(), &id)
// 	if err != nil{
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, publicvote)
// }

// func (h *Handler) GetAllPublicVotes(c *gin.Context){
// 	publicVotes, err := h.PublicVote.GetAllPublicVote(context.Background(), &g.PublicVote_Void{})
// 	if err != nil{
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, publicVotes)
// 	}

// func (h *Handler) UpdatePublicVote(c *gin.Context){
// 	publicvote := g.PublicVote{}
// 	err := c.BindJSON(&publicvote)
// 	if err != nil{
// 		c.JSON(400,  gin.H{"error": err.Error()})
// 		return
// 	}
// 	_, err = h.PublicVote.UpdatePublicVote(context.Background(), &g.UpdatePublicVoteRequest{PublicVote: &publicvote})
// 	if err != nil{
// 		c.JSON(500,  gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, gin.H{"message": "success"})
// }

// func (h *Handler) DeletePublicVote(c *gin.Context){
// 	id := g.DeletePublicVoteRequest{Id: c.Param(":id")}
// 	_, err := h.PublicVote.DeletePublicVote(context.Background(), &id)
// 	if err != nil{
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, gin.H{"message": "success"})
// }