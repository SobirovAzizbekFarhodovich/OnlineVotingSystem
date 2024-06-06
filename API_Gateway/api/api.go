package api

import (
	"github.com/GO/49-dars/API_Gateway/api/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/GO/49-dars/API_Gateway/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewGin(h *handler.Handler) *gin.Engine{
	r := gin.Default()
	r.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// r.GET("/elections", h.GetAllElections)
	// r.GET("/elections/:id", h.GetByIdElection)
	// r.DELETE("/election/:id", h.DeleteElection)
	// r.POST("/election", h.CreateElections)
	// r.PUT("/election", h.UpdateElection)


	// r.GET("/publicvotes", h.GetAllPublicVotes)
	// r.POST("/publicvote", h.CreatePublicVote)


	r.POST("/public", h.CreatePublic)
	r.GET("/public", h.GetAllPublic)
	r.GET("/public/:id", h.GetByIdPublic)
	r.DELETE("/public/:id", h.DeletePublic)
	r.PUT("/public", h.UpdatePublic)

	r.POST("/party",h.CreatePartys)
	r.GET("/party/:id",h.GetByIdParty)
	r.GET("/party",h.GetAllParty)
	r.PUT("/party",h.UpdateParty)
	r.DELETE("/party/:id",h.DeleteParty)



	return r
}