package api

import (
	"root/api/handler"
	_ "root/docs"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewGin(voting, public *grpc.ClientConn) *gin.Engine{
	h := handler.NewHandler(voting, public)

	r := gin.Default()
	r.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/elections", h.GetAllElections)
	r.GET("/election/:id", h.GetByIdElection)
	r.DELETE("/election/:id", h.DeleteElection)
	r.POST("/election", h.CreateElection)
	r.PUT("/election", h.UpdateElection)
	

	r.GET("/publicvotes", h.GetAllPublicVote) 
	r.POST("/publicvote", h.CreatePublicVote) 
	r.GET("/publicvote/:id", h.GetByIdPublicVote) 
	r.DELETE("/publicvote/:id", h.DeletePublicVote)
	r.PUT("/publicvote", h.UpdatePublicVote)


	r.POST("/candidate", h.CreateCandidate)
	r.GET("/candidate/:id", h.GetByIDCandidate)
	r.DELETE("candidate/:id", h.DeleteCandidate)
	r.PUT("/candidate", h.UpdateCandidate)
	r.GET("/candidates", h.GetAllCandidaties)

	r.POST("/party", h.CreateParty)
	r.GET("/parties", h.GetAllParty)
	r.GET("/party/:id", h.GetByIdParty)
	r.PUT("/party", h.UpdateParty)
	r.DELETE("/party/:id", h.DeleteParty)












	return r
}